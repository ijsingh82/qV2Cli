package node

import (
	"context"
	"math/big"

	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NodeCount() (*big.Int, error) {

	var d *big.Int
	client, err := ethclient.DialContext(context.Background(), cmd.Conf.Rpc)
	if err != nil {
		return d, err
	}
	defer client.Close()

	permImpl, err := b.NewNodeManager(common.HexToAddress(cmd.Conf.NodeMgrAddress), client)
	if err != nil {
		return d, err
	}

	d, err = permImpl.GetNumberOfNodes(nil)
	if err != nil {
		return d, err
	}
	return d, nil
}
