package impl

import (
	b "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"
	"github.com/ijsingh82/qV2Cli/pkg/model"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetPendingOp(orgId string) (model.PendingOp, error) {

	var op model.PendingOp
	// get client
	client, err := ethclient.Dial(cmd.Conf.Rpc)
	if err != nil {
		return op, err
	}
	defer client.Close()

	impl, err := b.NewPermImpl(common.HexToAddress(cmd.Conf.ImplAddress), client)
	if err != nil {
		return op, err
	}

	op.OrgId, op.EnodeId, op.Account, op.OpType, op.Err = impl.GetPendingOp(nil, orgId)
	if op.Err != nil {
		return op, op.Err
	}

	return op, nil
}
