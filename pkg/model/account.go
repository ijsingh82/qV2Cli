package model

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type AccountDetails struct {
	Id      common.Address
	OrgId   string
	Role    string
	Status  *big.Int
	IsAdmin bool
	Err     error
}

type AccountInfo struct {
	OrgId string
	Role  string
	Err   error
}

type Account struct {
	Status string
}

func (a *Account) AcctStatusString(stat *big.Int) (Account, error) {

	switch stat.Int64() {
	case 0:
		return Account{Status: "NotInList"}, nil
	case 1:
		return Account{Status: "AccountPendingApproval"}, nil
	case 2:
		return Account{Status: "Active"}, nil
	case 3:
		return Account{Status: "Inactive"}, nil
	case 4:
		return Account{Status: "Suspended"}, nil
	case 5:
		return Account{Status: "Blacklisted"}, nil
	case 6:
		return Account{Status: "Revoked"}, nil
	case 7:
		return Account{Status: "RecoveryInitiated"}, nil
	default:
		return Account{Status: ""}, errors.New("status not valid")
	}
}
