package impl

import (
	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"
	"github.com/ijsingh82/qV2Cli/pkg/model"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetPolicy() (model.PolicyDetails, error) {

	var pol model.PolicyDetails
	// get client
	client, err := ethclient.Dial(cmd.Conf.Rpc)
	if err != nil {
		return pol, err
	}
	defer client.Close()

	impl, err := b.NewPermImpl(common.HexToAddress(cmd.Conf.ImplAddress), client)
	if err != nil {
		return pol, err
	}

	pol.AdminOrg, pol.AdminRole, pol.OrgAdminRole, pol.NetworkBoot, err = impl.GetPolicyDetails(nil)
	if err != nil {
		return pol, err
	}

	return pol, nil
}
