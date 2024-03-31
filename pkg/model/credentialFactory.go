package model

import "github.com/ethereum/go-ethereum/common"

type CredentialFactory struct {
	Instituion  CredFacInst
	Name        string
	Symbol      string
	BaseUri     string
	Credentials []common.Address
	Holders     []common.Address
	TokenUri    []string
}

type CredFacInst struct {
	Approver        common.Address
	Name            string
	ContractAddress common.Address
}
