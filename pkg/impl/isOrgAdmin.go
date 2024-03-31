package impl

import (
	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func IsOrgAdmin(acctAdd, orgId string) (bool, error) {
	var orgAdmin bool

	// get client
	client, err := ethclient.Dial(cmd.Conf.Rpc)
	if err != nil {
		return orgAdmin, err
	}
	defer client.Close()

	impl, err := b.NewPermImpl(common.HexToAddress(cmd.Conf.ImplAddress), client)
	if err != nil {
		return orgAdmin, err
	}

	orgAdmin, err = impl.IsOrgAdmin(nil, common.HexToAddress(acctAdd), orgId)
	if err != nil {
		return orgAdmin, err
	}

	return orgAdmin, nil
}
