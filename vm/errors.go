package vm

import "errors"

var (
	ErrOutOfQuota                  = errors.New("out of quota")
	ErrDepth                       = errors.New("max call Depth exceeded")
	ErrInsufficientBalance         = errors.New("insufficient balance for transfer")
	ErrContractAddressCreationFail = errors.New("contract address creation fail")
	ErrAddressCollision            = errors.New("contract address collision")
	ErrTokenIdCreationFail         = errors.New("token id creation fail")
	ErrIdCollision                 = errors.New("id collision")
	ErrExecutionReverted           = errors.New("execution reverted")
	ErrInvalidData                 = errors.New("invalid data")
)

var (
	errGasUintOverflow       = errors.New("gas uint64 overflow")
	errReturnDataOutOfBounds = errors.New("evm: return data out of bounds")
)

var (
	ResultSuccess           = []byte{0x00}
	ResultErrorWithRetry    = []byte{0x01}
	ResultErrorWithoutRetry = []byte{0x02}
)