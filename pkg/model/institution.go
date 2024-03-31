package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Institution struct {
	Address     common.Address `json:"address,omitempty"`
	Approver    common.Address `json:"approver,omitempty"`
	Name        string         `json:"name,omitempty"`
	Uri         string         `json:"uri,omitempty"`
	Credentials []InstCred     `json:"credentials,omitempty"`
}

type InstCred struct {
	CredentialId   *big.Int       `json:"credential_id,omitempty"`
	InstApprover   common.Address `json:"inst_approver,omitempty"`
	CredApprover   common.Address `json:"cred_approver,omitempty"`
	Implementation common.Address `json:"implementation,omitempty"`
	CredAddress    common.Address `json:"cred_address,omitempty"`
	CredName       string         `json:"cred_name,omitempty"`
	CredSymbol     string         `json:"cred_symbol,omitempty"`
	CredUri        string         `json:"cred_uri,omitempty"`
}
