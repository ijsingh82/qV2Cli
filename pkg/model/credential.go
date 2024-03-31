package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Credential struct {
	Id                    *big.Int
	Status                *big.Int
	ExpirationDate        *big.Int
	CompletionDate        *big.Int
	InstitutionContract   common.Address
	AuthorizationContract common.Address
	Approver              common.Address
	Uri                   string
}
