package org

import (
	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"
	"github.com/ijsingh82/qV2Cli/pkg/model"

	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func FromID(orgId string) (model.OrgDetails, string, error) {

	var d model.OrgDetails
	var stat string

	client, err := ethclient.DialContext(context.Background(), cmd.Conf.Rpc)
	if err != nil {
		return d, stat, err
	}
	defer client.Close()

	orgMngr, err := b.NewOrgManager(common.HexToAddress(cmd.Conf.OrgMgrAddress), client)
	if err != nil {
		return d, stat, err
	}

	d.OrgId, d.PrntOrg, d.UltPrntOrg, d.Level, d.Status, err = orgMngr.GetOrgDetails(nil, orgId)
	if err != nil {
		return d, stat, err
	}

	stat, err = d.OrgStatusString()
	if err != nil {
		return d, stat, err
	}
	return d, stat, nil
}
