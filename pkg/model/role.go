package model

import (
	"errors"
	"math/big"
)

type RoleDetail struct {
	RoleId     string
	OrgId      string
	AccessType *big.Int
	Voter      bool
	Admin      bool
	Active     bool
}

func RoleAccesString(status *big.Int) (string, error) {

	switch status.Int64() {
	case 0:
		return "Read only", nil
	case 1:
		return "value transfer", nil
	case 2:
		return "contract deploy", nil
	case 3:
		return "full access", nil
	case 4:
		return "contract call", nil
	case 5:
		return "value transfer and contract call", nil
	case 6:
		return "value transfer and contract deploy", nil
	case 7:
		return "contract call and deploy", nil
	default:
		return "", errors.New("status not valid")
	}
}
