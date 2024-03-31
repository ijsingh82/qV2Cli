package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type DeployedContract struct {
	Name       string
	Address    common.Address
	Governance common.Address
}

type CredentialInfo struct {
	CredOrgAdd common.Address
	CredId     *big.Int
	Approver   common.Address
	Status     *big.Int
	Uri        string
}

// Config struct for the application
type Config struct {
	PermissionModel   string   `json:"PermissionModel"`
	UpgradableAddress string   `json:"upgradableAddress"`
	InterfaceAddress  string   `json:"interfaceAddress"`
	ImplAddress       string   `json:"implAddress"`
	NodeMgrAddress    string   `json:"nodeMgrAddress"`
	AccountMgrAddress string   `json:"accountMgrAddress"`
	RoleMgrAddress    string   `json:"roleMgrAddress"`
	VoterMgrAddress   string   `json:"voterMgrAddress"`
	OrgMgrAddress     string   `json:"orgMgrAddress"`
	NewAdminOrg       string   `json:"nwAdminOrg"`
	NewAdminRole      string   `json:"nwAdminRole"`
	OrgAdminRole      string   `json:"orgAdminRole"`
	GuardianAddress   []string `json:"accounts"`
	SubOrgDepth       int      `json:"subOrgDepth"`
	SubOrgBreadth     int      `json:"subOrgBreadth"`
}
