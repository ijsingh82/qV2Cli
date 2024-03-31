package model

import (
	"errors"
	"math/big"
)

type OrgDetails struct {
	OrgId      string
	PrntOrg    string
	UltPrntOrg string
	Level      *big.Int
	Status     *big.Int
}

func (o *OrgDetails) OrgStatusString() (string, error) {

	switch o.Status.Int64() {
	case 0:
		return "NotInList", nil
	case 1:
		return "Proposed", nil
	case 2:
		return "Approved", nil
	case 3:
		return "PendingSuspension", nil
	case 4:
		return "Suspended", nil
	case 5:
		return "AwaitingSuspensionRevoke", nil
	default:
		return "", errors.New("status not valid")
	}
}
