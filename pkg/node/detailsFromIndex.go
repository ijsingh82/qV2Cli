package node

import (
	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"
	"github.com/ijsingh82/qV2Cli/pkg/model"

	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// returns node details, status string and error
func FromIndex(index int64) (model.NodeDetails, string, error) {

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

	cnt, err := NodeCount()
	if err != nil {
		return d, "", err
	}

	if index > cnt.Int64()-1 {
		return d, "", fmt.Errorf("there are not that many nodes the current count is %v.  use an index less than count", cnt)
	}

	d, err = nodeMngr.GetNodeDetailsFromIndex(nil, big.NewInt(int64(index)))
	if err != nil {
		return d, "", err
	}

	stat, err := d.NodeStatusString()
	if err != nil {
		return d, "", err
	}

	return d, stat, nil
}
