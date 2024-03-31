package org

import (
	"math/big"

	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"

	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Count() (*big.Int, error) {

	var c *big.Int

	client, err := ethclient.DialContext(context.Background(), cmd.Conf.Rpc)
	if err != nil {
		return c, err
	}
	defer client.Close()

	orgMngr, err := b.NewOrgManager(common.HexToAddress(cmd.Conf.OrgMgrAddress), client)
	if err != nil {
		return c, err
	}

	c, err = orgMngr.GetNumberOfOrgs(nil)
	if err != nil {
		return c, err
	}

	return c, nil
}
