package pmchain

import (
	"fmt"
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/interfaces"
	"github.com/vitelabs/go-vite/ledger"
	"math/big"
)

func (c *chain) GetBalance(addr *types.Address, tokenId *types.TokenTypeId) (*big.Int, error) {
	return nil, nil
}

// get history balance, if history is too old, failed
func (c *chain) GetHistoryBalance(addr *types.Address, tokenId *types.TokenTypeId, accountBlockHash *types.Hash) (*big.Int, error) {
	return nil, nil
}

// get confirmed snapshot balance, if history is too old, failed
func (c *chain) GetConfirmedBalance(addr *types.Address, tokenId *types.TokenTypeId, sbHash *types.Hash) (*big.Int, error) {
	return nil, nil
}

// get contract code
func (c *chain) GetContractCode(contractAddr *types.Address) ([]byte, error) {
	return nil, nil
}

func (c *chain) GetContractMeta(contractAddress *types.Address) (meta *ledger.ContractMeta, err error) {
	return nil, nil
}

func (c *chain) GetContractList(gid *types.Gid) (map[types.Address]*ledger.ContractMeta, error) {
	return nil, nil
}

func (c *chain) GetStateSnapshot(blockHash *types.Hash) (interfaces.StateSnapshot, error) {
	stateSnapshot, err := c.stateDB.NewStateSnapshot(blockHash)
	if err != nil {
		c.log.Error(fmt.Sprintf("c.stateDB.NewStateSnapshot failed, error is %s, blockHash is %s", err, blockHash))
		return nil, err
	}
	return stateSnapshot, nil
}

func (c *chain) GetQuotaUnused(address *types.Address) uint64 {
	return 0
}

func (c *chain) GetQuotaUsed(address *types.Address) (quotaUsed uint64, blockCount uint64) {
	return 0, 0
}