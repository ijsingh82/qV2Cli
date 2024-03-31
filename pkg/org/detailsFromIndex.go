package org

import (
	"fmt"

	b "github.com/ijsingh82/qV2Cli/bind"

	"context"
	"math/big"

	"github.com/ijsingh82/qV2Cli/cmd"
	"github.com/ijsingh82/qV2Cli/pkg/model"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func FromIndex(index int64) (model.OrgDetails, string, error) {

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

	cnt, err := Count()
	if err != nil {
		return d, stat, err
	}

	if index > cnt.Int64()-1 {
		return d, stat, fmt.Errorf("the index %v is more than the count of orgs %v", index, cnt)
	}

	d.OrgId, d.PrntOrg, d.UltPrntOrg, d.Level, d.Status, err = orgMngr.GetOrgInfo(nil, big.NewInt(int64(index)))
	if err != nil {
		return d, stat, err
	}

	stat, err = d.OrgStatusString()
	if err != nil {
		return d, stat, err
	}

	return d, stat, err
}
