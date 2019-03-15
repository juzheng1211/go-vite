package chain_cache

import (
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/ledger"
)

type dataSet struct {
	dataId uint64

	dataRefCount map[uint64]int16

	blockDataId map[types.Hash]uint64

	accountBlockSet map[uint64]*ledger.AccountBlock

	snapshotBlockSet map[uint64]*ledger.SnapshotBlock

	abHeightIndexes map[types.Address]map[uint64]*ledger.AccountBlock

	sbHeightIndexes map[uint64]*ledger.SnapshotBlock
}

func NewDataSet() *dataSet {
	return &dataSet{
		dataRefCount: make(map[uint64]int16, 0),
		blockDataId:  make(map[types.Hash]uint64, 0),

		accountBlockSet:  make(map[uint64]*ledger.AccountBlock, 0),
		snapshotBlockSet: make(map[uint64]*ledger.SnapshotBlock, 0),

		abHeightIndexes: make(map[types.Address]map[uint64]*ledger.AccountBlock),

		sbHeightIndexes: make(map[uint64]*ledger.SnapshotBlock),
	}
}

func (ds *dataSet) RefDataId(dataId uint64) {
	if refCount, ok := ds.dataRefCount[dataId]; ok {
		ds.dataRefCount[dataId] = refCount + 1
	}
}
func (ds *dataSet) UnRefDataId(dataId uint64) {
	if refCount, ok := ds.dataRefCount[dataId]; ok {
		newRefCount := refCount - 1
		if newRefCount <= 0 {
			ds.gc(dataId)
		} else {
			ds.dataRefCount[dataId] = newRefCount
		}
	}
}

func (ds *dataSet) InsertAccountBlock(accountBlock *ledger.AccountBlock) uint64 {
	if dataId, ok := ds.blockDataId[accountBlock.Hash]; ok {
		return dataId
	}

	newDataId := ds.newDataId()

	// accountBlockSet
	ds.accountBlockSet[newDataId] = accountBlock

	// abHeightIndexes
	heightMap := ds.abHeightIndexes[accountBlock.AccountAddress]
	if heightMap == nil {
		heightMap = make(map[uint64]*ledger.AccountBlock)
	}
	heightMap[accountBlock.Height] = accountBlock
	ds.abHeightIndexes[accountBlock.AccountAddress] = heightMap

	// blockDataId
	ds.blockDataId[accountBlock.Hash] = newDataId

	return newDataId
}

func (ds *dataSet) InsertSnapshotBlock(snapshotBlock *ledger.SnapshotBlock) uint64 {
	if dataId, ok := ds.blockDataId[snapshotBlock.Hash]; ok {
		return dataId
	}

	newDataId := ds.newDataId()
	// snapshotBlockSet
	ds.snapshotBlockSet[newDataId] = snapshotBlock

	// sbHeightIndexes
	ds.sbHeightIndexes[snapshotBlock.Height] = snapshotBlock

	// blockDataId
	ds.blockDataId[snapshotBlock.Hash] = newDataId

	return newDataId
}

func (ds *dataSet) IsDataExisted(hash *types.Hash) bool {
	return ds.blockDataId[*hash] > 0
}

func (ds *dataSet) GetAccountBlock(dataId uint64) *ledger.AccountBlock {
	return ds.accountBlockSet[dataId]
}

func (ds *dataSet) GetSnapshotBlock(dataId uint64) *ledger.SnapshotBlock {
	return ds.snapshotBlockSet[dataId]
}

func (ds *dataSet) GetAccountBlockByHash(blockHash *types.Hash) *ledger.AccountBlock {
	dataId := ds.blockDataId[*blockHash]
	if dataId <= 0 {
		return nil
	}
	return ds.GetAccountBlock(dataId)
	//return ds.[*blockHash]
}

func (ds *dataSet) GetAccountBlockByHeight(address *types.Address, height uint64) *ledger.AccountBlock {
	abHeightMap := ds.abHeightIndexes[*address]
	if abHeightMap == nil {
		return nil
	}
	return abHeightMap[height]
}

func (ds *dataSet) GetSnapshotBlockByHash(blockHash *types.Hash) *ledger.SnapshotBlock {
	dataId := ds.blockDataId[*blockHash]
	if dataId <= 0 {
		return nil
	}
	return ds.GetSnapshotBlock(dataId)
}

func (ds *dataSet) GetSnapshotBlockByHeight(height uint64) *ledger.SnapshotBlock {
	return ds.sbHeightIndexes[height]
}

func (ds *dataSet) gc(dataId uint64) {

}

func (ds *dataSet) newDataId() uint64 {
	ds.dataId++
	return ds.dataId
}