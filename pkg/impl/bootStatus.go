package impl

import (
	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func BootStat() (bool, error) {
	var stat bool
	// get client
	client, err := ethclient.Dial(cmd.Conf.Rpc)
	if err != nil {
		return stat, err
	}
	defer client.Close()

	impl, err := b.NewPermImpl(common.HexToAddress(cmd.Conf.ImplAddress), client)
	if err != nil {
		return stat, err
	}

	stat, err = impl.GetNetworkBootStatus(nil)
	if err != nil {
		return stat, err
	}

	return stat, nil
}
