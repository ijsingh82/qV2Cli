package node

import (
	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"
	"github.com/ijsingh82/qV2Cli/pkg/model"

	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func FromEnode(enodeId string) (model.NodeDetails, string, error) {

	var d model.NodeDetails
	client, err := ethclient.DialContext(context.Background(), cmd.Conf.Rpc)
	if err != nil {
		return d, "", err
	}
	defer client.Close()

	nodeMngr, err := b.NewNodeManager(common.HexToAddress(cmd.Conf.NodeMgrAddress), client)
	if err != nil {
		return d, "", err
	}

	d, err = nodeMngr.GetNodeDetails(nil, enodeId)
	if err != nil {
		return d, "", err
	}

	stat, err := d.NodeStatusString()
	if err != nil {
		return d, "", err
	}
	return d, stat, nil
}
