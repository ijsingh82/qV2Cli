package impl

import (
	"math/big"

	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func IsTxAllowed(fromAdd, toAdd string, value int64, callData []byte) (bool, error) {

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

	allowed, err = impl.TransactionAllowed(nil, common.HexToAddress(fromAdd), common.HexToAddress(toAdd), big.NewInt(value), big.NewInt(0), big.NewInt(30000000), callData)
	if err != nil {
		return allowed, err
	}

	return allowed, nil
}
