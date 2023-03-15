package sequencer

import (
	"context"
	"math/big"
	"runtime"
	"sync"
	"time"

	"github.com/0xPolygonHermez/zkevm-node/log"
	"github.com/0xPolygonHermez/zkevm-node/state"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Worker represents the worker component of the sequencer
type Worker struct {
	cfg                  WorkerCfg
	pool                 map[string]*addrQueue
	efficiencyList       *efficiencyList
	workerMutex          sync.Mutex
	state                stateInterface
	batchConstraints     batchConstraints
	batchResourceWeights batchResourceWeights
}

// NewWorker creates an init a worker
func NewWorker(cfg WorkerCfg, state stateInterface, constraints batchConstraints, weights batchResourceWeights) *Worker {
	w := Worker{
		cfg:                  cfg,
		pool:                 make(map[string]*addrQueue),
		efficiencyList:       newEfficiencyList(),
		state:                state,
		batchConstraints:     constraints,
		batchResourceWeights: weights,
	}

	return &w
}

// NewTxTracker creates and inits a TxTracker
func (w *Worker) NewTxTracker(tx types.Transaction, isClaim bool, counters state.ZKCounters) (*TxTracker, error) {
	return newTxTracker(tx, isClaim, counters, w.batchConstraints, w.batchResourceWeights, w.cfg.ResourceCostMultiplier)
}

// AddTxTracker adds a new Tx to the Worker
func (w *Worker) AddTxTracker(ctx context.Context, tx *TxTracker) (dropTx bool) {
	w.workerMutex.Lock()
	defer w.workerMutex.Unlock()

	addr, found := w.pool[tx.FromStr]

	if !found {
		// Unlock the worker to let execute other worker functions while creating the new AddrQueue
		w.workerMutex.Unlock()

		root, err := w.state.GetLastStateRoot(ctx, nil)
		if err != nil {
			log.Errorf("AddTx GetLastStateRoot error: %v", err)
			return false
		}
		nonce, err := w.state.GetNonceByStateRoot(ctx, tx.From, root)
		if err != nil {
			log.Errorf("AddTx GetNonceByStateRoot error: %v", err)
			return false
		}
		balance, err := w.state.GetBalanceByStateRoot(ctx, tx.From, root)
		if err != nil {
			log.Errorf("AddTx GetBalanceByStateRoot error: %v", err)
			return false
		}

		addr = newAddrQueue(tx.From, nonce.Uint64(), balance)

		// Lock again the worker
		w.workerMutex.Lock()

		w.pool[tx.FromStr] = addr
		log.Infof("AddTx new addrQueue created for addr(%s)", tx.FromStr)
	}

	// Add the txTracker to Addr and get the newReadyTx and prevReadyTx
	log.Infof("AddTx new tx(%s) to addrQueue(%s)", tx.Hash.String(), tx.FromStr)
	newReadyTx, prevReadyTx, dropTx := addr.addTx(tx)
	if dropTx {
		log.Infof("AddTx tx(%s) dropped from addrQueue(%s)", tx.Hash.String(), tx.FromStr)
		return dropTx
	}

	// Update the EfficiencyList (if needed)
	if prevReadyTx != nil {
		log.Infof("AddTx prevReadyTx(%s) deleted from EfficiencyList", prevReadyTx.Hash.String())
		w.efficiencyList.delete(prevReadyTx)
	}
	if newReadyTx != nil {
		log.Infof("AddTx newReadyTx(%s) added to EfficiencyList", newReadyTx.Hash.String())
		w.efficiencyList.add(newReadyTx)
	}

	return false
}

func (w *Worker) applyAddressUpdate(from common.Address, fromNonce *uint64, fromBalance *big.Int) (*TxTracker, *TxTracker, []*TxTracker) {
	addrQueue, found := w.pool[from.String()]

	if found {
		newReadyTx, prevReadyTx, txsToDelete := addrQueue.updateCurrentNonceBalance(fromNonce, fromBalance)

		// Update the EfficiencyList (if needed)
		if prevReadyTx != nil {
			log.Infof("applyAddressUpdate prevReadyTx(%s) deleted from EfficiencyList", prevReadyTx.Hash.String())
			w.efficiencyList.delete(prevReadyTx)
		}
		if newReadyTx != nil {
			log.Infof("applyAddressUpdate newReadyTx(%s) added to EfficiencyList", newReadyTx.Hash.String())
			w.efficiencyList.add(newReadyTx)
		}

		return newReadyTx, prevReadyTx, txsToDelete
	}

	return nil, nil, nil
}

// UpdateAfterSingleSuccessfulTxExecution updates the touched addresses after execute on Executor a successfully tx
func (w *Worker) UpdateAfterSingleSuccessfulTxExecution(from common.Address, touchedAddresses map[common.Address]*state.InfoReadWrite) []*TxTracker {
	w.workerMutex.Lock()
	defer w.workerMutex.Unlock()
	if len(touchedAddresses) == 0 {
		log.Errorf("UpdateAfterSingleSuccessfulTxExecution touchedAddresses is nil or empty")
	}
	txsToDelete := make([]*TxTracker, 0)
	touchedFrom, found := touchedAddresses[from]
	if found {
		fromNonce, fromBalance := touchedFrom.Nonce, touchedFrom.Balance
		_, _, txsToDelete = w.applyAddressUpdate(from, fromNonce, fromBalance)
	} else {
		log.Errorf("UpdateAfterSingleSuccessfulTxExecution from(%s) not found in touchedAddresses", from.String())
	}

	for addr, addressInfo := range touchedAddresses {
		if addr != from {
			_, _, txsToDeleteTemp := w.applyAddressUpdate(addr, nil, addressInfo.Balance)
			txsToDelete = append(txsToDelete, txsToDeleteTemp...)
		}
	}
	return txsToDelete
}

// MoveTxToNotReady move a tx to not ready after it fails to execute
func (w *Worker) MoveTxToNotReady(txHash common.Hash, from common.Address, actualNonce *uint64, actualBalance *big.Int) []*TxTracker {
	w.workerMutex.Lock()
	defer w.workerMutex.Unlock()

	addrQueue, found := w.pool[from.String()]
	if found {
		// Sanity check. The txHash must be the readyTx
		if addrQueue.readyTx == nil || txHash.String() != addrQueue.readyTx.HashStr {
			readyHashStr := ""
			if addrQueue.readyTx != nil {
				readyHashStr = addrQueue.readyTx.HashStr
			}
			log.Errorf("MoveTxToNotReady txHash(s) is not the readyTx(%s)", txHash.String(), readyHashStr)
		}
	}
	_, _, txsToDelete := w.applyAddressUpdate(from, actualNonce, actualBalance)

	return txsToDelete
}

// DeleteTx delete the tx after it fails to execute
func (w *Worker) DeleteTx(txHash common.Hash, addr common.Address) {
	w.workerMutex.Lock()
	defer w.workerMutex.Unlock()

	addrQueue, found := w.pool[addr.String()]
	if found {
		deletedReadyTx := addrQueue.deleteTx(txHash)
		if deletedReadyTx != nil {
			log.Infof("DeleteTx tx(%s) deleted from EfficiencyList", deletedReadyTx.Hash.String())
			w.efficiencyList.delete(deletedReadyTx)
		}
	} else {
		log.Errorf("DeleteTx addrQueue(%s) not found", addr.String())
	}
}

// UpdateTx updates the ZKCounter of a tx and resort the tx in the efficiency list if needed
func (w *Worker) UpdateTx(txHash common.Hash, addr common.Address, counters state.ZKCounters) {
	w.workerMutex.Lock()
	defer w.workerMutex.Unlock()

	addrQueue, found := w.pool[addr.String()]

	if found {
		newReadyTx, prevReadyTx := addrQueue.UpdateTxZKCounters(txHash, counters)

		// Resort the newReadyTx in efficiencyList
		if prevReadyTx != nil {
			log.Infof("UpdateTx prevReadyTx(%s) deleted from EfficiencyList", prevReadyTx.Hash.String())
			w.efficiencyList.delete(prevReadyTx)
		}
		if newReadyTx != nil {
			log.Infof("UpdateTx newReadyTx(%s) added to EfficiencyList", newReadyTx.Hash.String())
			w.efficiencyList.add(newReadyTx)
		}
	} else {
		log.Errorf("UpdateTx addrQueue(%s) not found", addr.String())
	}
}

// GetBestFittingTx gets the most efficient tx that fits in the available batch resources
func (w *Worker) GetBestFittingTx(resources batchResources) *TxTracker {
	w.workerMutex.Lock()
	defer w.workerMutex.Unlock()

	var (
		tx         *TxTracker
		foundMutex sync.RWMutex
	)

	nGoRoutines := runtime.NumCPU()
	foundAt := -1

	wg := sync.WaitGroup{}
	wg.Add(nGoRoutines)

	// Each go routine looks for a fitting tx
	for i := 0; i < nGoRoutines; i++ {
		go func(n int, bresources batchResources) {
			defer wg.Done()
			for i := n; i < w.efficiencyList.len(); i += nGoRoutines {
				foundMutex.RLock()
				if foundAt != -1 && i > foundAt {
					foundMutex.RUnlock()
					return
				}
				foundMutex.RUnlock()

				txCandidate := w.efficiencyList.getByIndex(i)
				err := bresources.sub(txCandidate.BatchResources)
				if err != nil {
					// We don't add this Tx
					continue
				}

				foundMutex.Lock()
				if foundAt == -1 || foundAt > i {
					foundAt = i
					tx = txCandidate
					log.Infof("GetBestFittingTx found tx(%s) at index(%d)", tx.Hash.String(), i)
				}
				foundMutex.Unlock()

				return
			}
		}(i, resources)
	}
	wg.Wait()

	return tx
}

// ExpireTransactions deletes old txs
func (w *Worker) ExpireTransactions(maxTime time.Duration) []*TxTracker {
	w.workerMutex.Lock()
	defer w.workerMutex.Unlock()

	var txs []*TxTracker

	log.Info("ExpireTransactions start. addrQueue len: ", len(w.pool))
	for _, addrQueue := range w.pool {
		subTxs, prevReadyTx := addrQueue.ExpireTransactions(maxTime)
		txs = append(txs, subTxs...)

		if prevReadyTx != nil {
			w.efficiencyList.delete(prevReadyTx)
		}

		if addrQueue.IsEmpty() {
			delete(w.pool, addrQueue.fromStr)
		}
	}
	log.Info("ExpireTransactions end. addrQueue len: ", len(w.pool), "deleteCount: ", len(txs))

	return txs
}

// GetEfficiencyList returns the efficiency list
func (w *Worker) GetEfficiencyList() *efficiencyList {
	return w.efficiencyList
}

// HandleL2Reorg handles the L2 reorg signal
func (w *Worker) HandleL2Reorg(txHashes []common.Hash) {
	log.Fatal("L2 Reorg detected. Restarting to sync with the new L2 state...")
}
