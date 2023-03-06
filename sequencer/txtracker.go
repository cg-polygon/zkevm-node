package sequencer

import (
	"math"
	"math/big"
	"time"

	"github.com/0xPolygonHermez/zkevm-node/log"
	"github.com/0xPolygonHermez/zkevm-node/state"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TxTracker is a struct that contains all the tx data needed to be managed by the worker
type TxTracker struct {
	Hash                   common.Hash
	HashStr                string
	From                   common.Address
	FromStr                string
	Nonce                  uint64
	Gas                    uint64 // To check if it fits into a batch
	GasPrice               *big.Int
	Cost                   *big.Int       // Cost = Amount + Benefit
	Benefit                *big.Int       // GasLimit * GasPrice
	IsClaim                bool           // Needed to calculate efficiency
	BatchResources         batchResources // To check if it fits into a batch
	Efficiency             float64
	RawTx                  []byte
	constraints            batchConstraintsFloat64
	weightMultipliers      batchResourceWeightMultipliers
	resourceCostMultiplier float64
	totalWeight            float64
	ReceivedAt             time.Time // To
}

// batchResourceWeightMultipliers is a struct that contains the weight multipliers for each resource
type batchResourceWeightMultipliers struct {
	cumulativeGasUsed float64
	arithmetics       float64
	binaries          float64
	keccakHashes      float64
	memAligns         float64
	poseidonHashes    float64
	poseidonPaddings  float64
	steps             float64
	batchBytesSize    float64
}

// batchConstraints represents the constraints for a batch in float64
type batchConstraintsFloat64 struct {
	maxTxsPerBatch       float64
	maxBatchBytesSize    float64
	maxCumulativeGasUsed float64
	maxKeccakHashes      float64
	maxPoseidonHashes    float64
	maxPoseidonPaddings  float64
	maxMemAligns         float64
	maxArithmetics       float64
	maxBinaries          float64
	maxSteps             float64
}

// newTxTracker creates and inti a TxTracker
func newTxTracker(tx types.Transaction, isClaim bool, counters state.ZKCounters, constraints batchConstraints, weights batchResourceWeights, resourceCostMultiplier float64) (*TxTracker, error) {
	addr, err := state.GetSender(tx)
	if err != nil {
		return nil, err
	}

	txTracker := &TxTracker{
		IsClaim:  isClaim,
		Hash:     tx.Hash(),
		HashStr:  tx.Hash().String(),
		From:     addr,
		FromStr:  addr.String(),
		Nonce:    tx.Nonce(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Cost:     tx.Cost(),
		BatchResources: batchResources{
			bytes:      tx.Size(),
			zKCounters: counters,
		},
		resourceCostMultiplier: resourceCostMultiplier,
		ReceivedAt:             time.Now(),
	}

	txTracker.Benefit = new(big.Int).Mul(new(big.Int).SetUint64(txTracker.Gas), txTracker.GasPrice)
	txTracker.constraints = convertBatchConstraintsToFloat64(constraints)
	txTracker.totalWeight = float64(weights.WeightArithmetics + weights.WeightBatchBytesSize + weights.WeightBinaries + weights.WeightCumulativeGasUsed +
		weights.WeightKeccakHashes + weights.WeightMemAligns + weights.WeightPoseidonHashes + weights.WeightPoseidonPaddings + weights.WeightSteps)
	txTracker.weightMultipliers = calculateWeightMultipliers(weights, txTracker.totalWeight)
	txTracker.calculateEfficiency()
	txTracker.RawTx, err = state.EncodeTransactions([]types.Transaction{tx})
	if err != nil {
		return nil, err
	}

	return txTracker, nil
}

// updateZKCounters updates the counters of the tx and recalculates the tx efficiency
func (tx *TxTracker) updateZKCounters(counters state.ZKCounters) {
	tx.BatchResources.zKCounters = counters
	tx.calculateEfficiency()
}

// calculateEfficiency calculates the tx efficiency
func (tx *TxTracker) calculateEfficiency() {
	resourceCost := (float64(tx.BatchResources.zKCounters.CumulativeGasUsed)/tx.constraints.maxCumulativeGasUsed)*tx.weightMultipliers.cumulativeGasUsed +
		(float64(tx.BatchResources.zKCounters.UsedArithmetics)/tx.constraints.maxArithmetics)*tx.weightMultipliers.arithmetics +
		(float64(tx.BatchResources.zKCounters.UsedBinaries)/tx.constraints.maxBinaries)*tx.weightMultipliers.binaries +
		(float64(tx.BatchResources.zKCounters.UsedKeccakHashes)/tx.constraints.maxKeccakHashes)*tx.weightMultipliers.keccakHashes +
		(float64(tx.BatchResources.zKCounters.UsedMemAligns)/tx.constraints.maxMemAligns)*tx.weightMultipliers.memAligns +
		(float64(tx.BatchResources.zKCounters.UsedPoseidonHashes)/tx.constraints.maxPoseidonHashes)*tx.weightMultipliers.poseidonHashes +
		(float64(tx.BatchResources.zKCounters.UsedPoseidonPaddings)/tx.constraints.maxPoseidonPaddings)*tx.weightMultipliers.poseidonPaddings +
		(float64(tx.BatchResources.zKCounters.UsedSteps)/tx.constraints.maxSteps)*tx.weightMultipliers.steps +
		(float64(tx.BatchResources.bytes)/tx.constraints.maxBatchBytesSize)*tx.weightMultipliers.batchBytesSize

	resourceCost = resourceCost * tx.resourceCostMultiplier
	var eff *big.Float
	if tx.IsClaim {
		eff = big.NewFloat(math.MaxFloat64)
	} else {
		ben := big.NewFloat(0).SetInt(tx.Benefit)
		rc := big.NewFloat(0).SetFloat64(resourceCost)
		eff = big.NewFloat(0).Quo(ben, rc)
	}

	var accuracy big.Accuracy
	tx.Efficiency, accuracy = eff.Float64()
	if accuracy != big.Exact {
		log.Errorf("CalculateEfficiency accuracy warning (%s). Calculated=%s Assigned=%f", accuracy.String(), eff.String(), tx.Efficiency)
	}
}

// convertBatchConstraintsToFloat64 converts the batch constraints to float64
func convertBatchConstraintsToFloat64(constraints batchConstraints) batchConstraintsFloat64 {
	return batchConstraintsFloat64{
		maxTxsPerBatch:       float64(constraints.MaxTxsPerBatch),
		maxBatchBytesSize:    float64(constraints.MaxBatchBytesSize),
		maxCumulativeGasUsed: float64(constraints.MaxCumulativeGasUsed),
		maxKeccakHashes:      float64(constraints.MaxKeccakHashes),
		maxPoseidonHashes:    float64(constraints.MaxPoseidonHashes),
		maxPoseidonPaddings:  float64(constraints.MaxPoseidonPaddings),
		maxMemAligns:         float64(constraints.MaxMemAligns),
		maxArithmetics:       float64(constraints.MaxArithmetics),
		maxBinaries:          float64(constraints.MaxBinaries),
		maxSteps:             float64(constraints.MaxSteps),
	}
}

// calculateWeightMultipliers calculates the weight multipliers for each resource
func calculateWeightMultipliers(weights batchResourceWeights, totalWeight float64) batchResourceWeightMultipliers {
	return batchResourceWeightMultipliers{
		cumulativeGasUsed: float64(weights.WeightCumulativeGasUsed) / totalWeight,
		arithmetics:       float64(weights.WeightArithmetics) / totalWeight,
		binaries:          float64(weights.WeightBinaries) / totalWeight,
		keccakHashes:      float64(weights.WeightKeccakHashes) / totalWeight,
		memAligns:         float64(weights.WeightMemAligns) / totalWeight,
		poseidonHashes:    float64(weights.WeightPoseidonHashes) / totalWeight,
		poseidonPaddings:  float64(weights.WeightPoseidonPaddings) / totalWeight,
		steps:             float64(weights.WeightSteps) / totalWeight,
		batchBytesSize:    float64(weights.WeightBatchBytesSize) / totalWeight,
	}
}
