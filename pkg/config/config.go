package config

import (
	"encoding/json"
	"fmt"
)

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
	GuardianAddress   []string `json:"accounts"`
	NewAdminOrg       string   `json:"nwAdminOrg"`
	NwAdminRole       string   `json:"nwAdminRole"`
	OrgAdminRole      string   `json:"orgAdminRole"`
	SubOrgDepth       int      `json:"subOrgDepth"`
	SubOrgBreadth     int      `json:"subOrgBreadth"`
	Rpc               string   `json:"rpc"`
}

// Load loads the configuration from the specified file
func Load(configFile1 []byte) (*Config, error) {
	conf := &Config{}
	if err := json.Unmarshal(configFile1, conf); err != nil {
		fmt.Println("Failed to parse config file: ", err)
	}

	return conf, nil
}
