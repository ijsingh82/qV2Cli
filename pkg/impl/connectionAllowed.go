package impl

import (
	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ConnAllowed(enodeID, ip string, port uint16) (bool, error) {

	var allowed bool
	// get client
	client, err := ethclient.Dial(cmd.Conf.Rpc)
	if err != nil {
		return allowed, err
	}
	defer client.Close()

	impl, err := b.NewPermImpl(common.HexToAddress(cmd.Conf.ImplAddress), client)
	if err != nil {
		return allowed, err
	}

	allowed, err = impl.ConnectionAllowed(nil, enodeID, ip, port)
	if err != nil {
		return allowed, err
	}

	return allowed, nil
}
