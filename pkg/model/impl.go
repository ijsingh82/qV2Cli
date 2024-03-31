package model

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type PendingOp struct {
	OrgId   string
	EnodeId string
	Account common.Address
	OpType  *big.Int
	Err     error
}

type PolicyDetails struct {
	AdminOrg     string
	AdminRole    string
	OrgAdminRole string
	NetworkBoot  bool
}

// 0 - None - indicates no pending records for the org
// 1 - New org add activity
// 2 - Org suspension activity
// 3 - Revoke of org suspension
// 4 - Assigning admin role for a new account
// 5 - Blacklisted node recovery
// 6 - Blacklisted account recovery
func (p *PendingOp) PendingOpStatusString() (Account, error) {

	switch p.OpType.Int64() {
	case 0:
		return Account{Status: "None - indicates no pending records for the org"}, nil
	case 1:
		return Account{Status: "New org add activity"}, nil
	case 2:
		return Account{Status: "Org suspension activity"}, nil
	case 3:
		return Account{Status: "Revoke of org suspension"}, nil
	case 4:
		return Account{Status: "Assigning admin role for a new account"}, nil
	case 5:
		return Account{Status: "Blacklisted node recovery"}, nil
	case 6:
		return Account{Status: "Blacklisted account recovery"}, nil
	default:
		return Account{Status: ""}, errors.New("status not valid")
	}
}
