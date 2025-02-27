// Code generated by mockery v2.16.0. DO NOT EDIT.

package jsonrpc

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v4"

	runtime "github.com/0xPolygonHermez/zkevm-node/state/runtime"

	state "github.com/0xPolygonHermez/zkevm-node/state"

	time "time"

	types "github.com/ethereum/go-ethereum/core/types"
)

// stateMock is an autogenerated mock type for the stateInterface type
type stateMock struct {
	mock.Mock
}

// BatchNumberByL2BlockNumber provides a mock function with given fields: ctx, blockNumber, dbTx
func (_m *stateMock) BatchNumberByL2BlockNumber(ctx context.Context, blockNumber uint64, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, blockNumber, dbTx)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) uint64); ok {
		r0 = rf(ctx, blockNumber, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, blockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BeginStateTransaction provides a mock function with given fields: ctx
func (_m *stateMock) BeginStateTransaction(ctx context.Context) (pgx.Tx, error) {
	ret := _m.Called(ctx)

	var r0 pgx.Tx
	if rf, ok := ret.Get(0).(func(context.Context) pgx.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Tx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DebugTransaction provides a mock function with given fields: ctx, transactionHash, traceConfig, dbTx
func (_m *stateMock) DebugTransaction(ctx context.Context, transactionHash common.Hash, traceConfig state.TraceConfig, dbTx pgx.Tx) (*runtime.ExecutionResult, error) {
	ret := _m.Called(ctx, transactionHash, traceConfig, dbTx)

	var r0 *runtime.ExecutionResult
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, state.TraceConfig, pgx.Tx) *runtime.ExecutionResult); ok {
		r0 = rf(ctx, transactionHash, traceConfig, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*runtime.ExecutionResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, state.TraceConfig, pgx.Tx) error); ok {
		r1 = rf(ctx, transactionHash, traceConfig, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EstimateGas provides a mock function with given fields: transaction, senderAddress, l2BlockNumber, dbTx
func (_m *stateMock) EstimateGas(transaction *types.Transaction, senderAddress common.Address, l2BlockNumber *uint64, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(transaction, senderAddress, l2BlockNumber, dbTx)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*types.Transaction, common.Address, *uint64, pgx.Tx) uint64); ok {
		r0 = rf(transaction, senderAddress, l2BlockNumber, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*types.Transaction, common.Address, *uint64, pgx.Tx) error); ok {
		r1 = rf(transaction, senderAddress, l2BlockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBalance provides a mock function with given fields: ctx, address, blockNumber, dbTx
func (_m *stateMock) GetBalance(ctx context.Context, address common.Address, blockNumber uint64, dbTx pgx.Tx) (*big.Int, error) {
	ret := _m.Called(ctx, address, blockNumber, dbTx)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, uint64, pgx.Tx) *big.Int); ok {
		r0 = rf(ctx, address, blockNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, address, blockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBatchByNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *stateMock) GetBatchByNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (*state.Batch, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 *state.Batch
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *state.Batch); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Batch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCode provides a mock function with given fields: ctx, address, blockNumber, dbTx
func (_m *stateMock) GetCode(ctx context.Context, address common.Address, blockNumber uint64, dbTx pgx.Tx) ([]byte, error) {
	ret := _m.Called(ctx, address, blockNumber, dbTx)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, uint64, pgx.Tx) []byte); ok {
		r0 = rf(ctx, address, blockNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, address, blockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetL2BlockByHash provides a mock function with given fields: ctx, hash, dbTx
func (_m *stateMock) GetL2BlockByHash(ctx context.Context, hash common.Hash, dbTx pgx.Tx) (*types.Block, error) {
	ret := _m.Called(ctx, hash, dbTx)

	var r0 *types.Block
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, pgx.Tx) *types.Block); ok {
		r0 = rf(ctx, hash, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, pgx.Tx) error); ok {
		r1 = rf(ctx, hash, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetL2BlockByNumber provides a mock function with given fields: ctx, blockNumber, dbTx
func (_m *stateMock) GetL2BlockByNumber(ctx context.Context, blockNumber uint64, dbTx pgx.Tx) (*types.Block, error) {
	ret := _m.Called(ctx, blockNumber, dbTx)

	var r0 *types.Block
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *types.Block); ok {
		r0 = rf(ctx, blockNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, blockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetL2BlockHashesSince provides a mock function with given fields: ctx, since, dbTx
func (_m *stateMock) GetL2BlockHashesSince(ctx context.Context, since time.Time, dbTx pgx.Tx) ([]common.Hash, error) {
	ret := _m.Called(ctx, since, dbTx)

	var r0 []common.Hash
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, pgx.Tx) []common.Hash); ok {
		r0 = rf(ctx, since, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, time.Time, pgx.Tx) error); ok {
		r1 = rf(ctx, since, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetL2BlockHeaderByNumber provides a mock function with given fields: ctx, blockNumber, dbTx
func (_m *stateMock) GetL2BlockHeaderByNumber(ctx context.Context, blockNumber uint64, dbTx pgx.Tx) (*types.Header, error) {
	ret := _m.Called(ctx, blockNumber, dbTx)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *types.Header); ok {
		r0 = rf(ctx, blockNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, blockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetL2BlockTransactionCountByHash provides a mock function with given fields: ctx, hash, dbTx
func (_m *stateMock) GetL2BlockTransactionCountByHash(ctx context.Context, hash common.Hash, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, hash, dbTx)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, pgx.Tx) uint64); ok {
		r0 = rf(ctx, hash, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, pgx.Tx) error); ok {
		r1 = rf(ctx, hash, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetL2BlockTransactionCountByNumber provides a mock function with given fields: ctx, blockNumber, dbTx
func (_m *stateMock) GetL2BlockTransactionCountByNumber(ctx context.Context, blockNumber uint64, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, blockNumber, dbTx)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) uint64); ok {
		r0 = rf(ctx, blockNumber, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, blockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastBatchNumber provides a mock function with given fields: ctx, dbTx
func (_m *stateMock) GetLastBatchNumber(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) uint64); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastConsolidatedL2BlockNumber provides a mock function with given fields: ctx, dbTx
func (_m *stateMock) GetLastConsolidatedL2BlockNumber(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) uint64); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastL2Block provides a mock function with given fields: ctx, dbTx
func (_m *stateMock) GetLastL2Block(ctx context.Context, dbTx pgx.Tx) (*types.Block, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 *types.Block
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) *types.Block); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastL2BlockNumber provides a mock function with given fields: ctx, dbTx
func (_m *stateMock) GetLastL2BlockNumber(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) uint64); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastVerifiedBatch provides a mock function with given fields: ctx, dbTx
func (_m *stateMock) GetLastVerifiedBatch(ctx context.Context, dbTx pgx.Tx) (*state.VerifiedBatch, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 *state.VerifiedBatch
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) *state.VerifiedBatch); ok {
		r0 = rf(ctx, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.VerifiedBatch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastVirtualBatchNum provides a mock function with given fields: ctx, dbTx
func (_m *stateMock) GetLastVirtualBatchNum(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) uint64); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLogs provides a mock function with given fields: ctx, fromBlock, toBlock, addresses, topics, blockHash, since, dbTx
func (_m *stateMock) GetLogs(ctx context.Context, fromBlock uint64, toBlock uint64, addresses []common.Address, topics [][]common.Hash, blockHash *common.Hash, since *time.Time, dbTx pgx.Tx) ([]*types.Log, error) {
	ret := _m.Called(ctx, fromBlock, toBlock, addresses, topics, blockHash, since, dbTx)

	var r0 []*types.Log
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, []common.Address, [][]common.Hash, *common.Hash, *time.Time, pgx.Tx) []*types.Log); ok {
		r0 = rf(ctx, fromBlock, toBlock, addresses, topics, blockHash, since, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Log)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, []common.Address, [][]common.Hash, *common.Hash, *time.Time, pgx.Tx) error); ok {
		r1 = rf(ctx, fromBlock, toBlock, addresses, topics, blockHash, since, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNonce provides a mock function with given fields: ctx, address, blockNumber, dbTx
func (_m *stateMock) GetNonce(ctx context.Context, address common.Address, blockNumber uint64, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, address, blockNumber, dbTx)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, uint64, pgx.Tx) uint64); ok {
		r0 = rf(ctx, address, blockNumber, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, address, blockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageAt provides a mock function with given fields: ctx, address, position, blockNumber, dbTx
func (_m *stateMock) GetStorageAt(ctx context.Context, address common.Address, position *big.Int, blockNumber uint64, dbTx pgx.Tx) (*big.Int, error) {
	ret := _m.Called(ctx, address, position, blockNumber, dbTx)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int, uint64, pgx.Tx) *big.Int); ok {
		r0 = rf(ctx, address, position, blockNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, address, position, blockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSyncingInfo provides a mock function with given fields: ctx, dbTx
func (_m *stateMock) GetSyncingInfo(ctx context.Context, dbTx pgx.Tx) (state.SyncingInfo, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 state.SyncingInfo
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) state.SyncingInfo); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(state.SyncingInfo)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionByHash provides a mock function with given fields: ctx, transactionHash, dbTx
func (_m *stateMock) GetTransactionByHash(ctx context.Context, transactionHash common.Hash, dbTx pgx.Tx) (*types.Transaction, error) {
	ret := _m.Called(ctx, transactionHash, dbTx)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, pgx.Tx) *types.Transaction); ok {
		r0 = rf(ctx, transactionHash, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, pgx.Tx) error); ok {
		r1 = rf(ctx, transactionHash, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionByL2BlockHashAndIndex provides a mock function with given fields: ctx, blockHash, index, dbTx
func (_m *stateMock) GetTransactionByL2BlockHashAndIndex(ctx context.Context, blockHash common.Hash, index uint64, dbTx pgx.Tx) (*types.Transaction, error) {
	ret := _m.Called(ctx, blockHash, index, dbTx)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, uint64, pgx.Tx) *types.Transaction); ok {
		r0 = rf(ctx, blockHash, index, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, blockHash, index, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionByL2BlockNumberAndIndex provides a mock function with given fields: ctx, blockNumber, index, dbTx
func (_m *stateMock) GetTransactionByL2BlockNumberAndIndex(ctx context.Context, blockNumber uint64, index uint64, dbTx pgx.Tx) (*types.Transaction, error) {
	ret := _m.Called(ctx, blockNumber, index, dbTx)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, pgx.Tx) *types.Transaction); ok {
		r0 = rf(ctx, blockNumber, index, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, blockNumber, index, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionReceipt provides a mock function with given fields: ctx, transactionHash, dbTx
func (_m *stateMock) GetTransactionReceipt(ctx context.Context, transactionHash common.Hash, dbTx pgx.Tx) (*types.Receipt, error) {
	ret := _m.Called(ctx, transactionHash, dbTx)

	var r0 *types.Receipt
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, pgx.Tx) *types.Receipt); ok {
		r0 = rf(ctx, transactionHash, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Receipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, pgx.Tx) error); ok {
		r1 = rf(ctx, transactionHash, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionsByBatchNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *stateMock) GetTransactionsByBatchNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) ([]types.Transaction, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 []types.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) []types.Transaction); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVerifiedBatch provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *stateMock) GetVerifiedBatch(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (*state.VerifiedBatch, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 *state.VerifiedBatch
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *state.VerifiedBatch); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.VerifiedBatch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVirtualBatch provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *stateMock) GetVirtualBatch(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (*state.VirtualBatch, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 *state.VirtualBatch
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *state.VirtualBatch); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.VirtualBatch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsL2BlockConsolidated provides a mock function with given fields: ctx, blockNumber, dbTx
func (_m *stateMock) IsL2BlockConsolidated(ctx context.Context, blockNumber uint64, dbTx pgx.Tx) (bool, error) {
	ret := _m.Called(ctx, blockNumber, dbTx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) bool); ok {
		r0 = rf(ctx, blockNumber, dbTx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, blockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsL2BlockVirtualized provides a mock function with given fields: ctx, blockNumber, dbTx
func (_m *stateMock) IsL2BlockVirtualized(ctx context.Context, blockNumber uint64, dbTx pgx.Tx) (bool, error) {
	ret := _m.Called(ctx, blockNumber, dbTx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) bool); ok {
		r0 = rf(ctx, blockNumber, dbTx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, blockNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PrepareWebSocket provides a mock function with given fields:
func (_m *stateMock) PrepareWebSocket() {
	_m.Called()
}

// ProcessUnsignedTransaction provides a mock function with given fields: ctx, tx, senderAddress, l2BlockNumber, noZKEVMCounters, dbTx
func (_m *stateMock) ProcessUnsignedTransaction(ctx context.Context, tx *types.Transaction, senderAddress common.Address, l2BlockNumber uint64, noZKEVMCounters bool, dbTx pgx.Tx) *runtime.ExecutionResult {
	ret := _m.Called(ctx, tx, senderAddress, l2BlockNumber, noZKEVMCounters, dbTx)

	var r0 *runtime.ExecutionResult
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction, common.Address, uint64, bool, pgx.Tx) *runtime.ExecutionResult); ok {
		r0 = rf(ctx, tx, senderAddress, l2BlockNumber, noZKEVMCounters, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*runtime.ExecutionResult)
		}
	}

	return r0
}

// RegisterNewL2BlockEventHandler provides a mock function with given fields: h
func (_m *stateMock) RegisterNewL2BlockEventHandler(h state.NewL2BlockEventHandler) {
	_m.Called(h)
}

type mockConstructorTestingTnewStateMock interface {
	mock.TestingT
	Cleanup(func())
}

// newStateMock creates a new instance of stateMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newStateMock(t mockConstructorTestingTnewStateMock) *stateMock {
	mock := &stateMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
