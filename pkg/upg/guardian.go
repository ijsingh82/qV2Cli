package upg

import (
	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"

	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Guardian() (common.Address, error) {

	var add common.Address
	client, err := ethclient.DialContext(context.Background(), cmd.Conf.Rpc)
	if err != nil {
		return add, err
	}
	defer client.Close()

	permUpg, err := b.NewPermUpgr(common.HexToAddress(cmd.Conf.UpgradableAddress), client)
	if err != nil {
		return add, err
	}

	add, err = permUpg.GetGuardian(nil)
	if err != nil {
		return add, err
	}

	return add, nil
}
