package model

import (
	"errors"
	"math/big"
)

type CreatedNode struct {
	ValidatorAccount string
	NodePath         string
}

type NodeDetails struct {
	OrgId      string
	EnodeId    string
	Ip         string
	Port       uint16
	Raftport   uint16
	NodeStatus *big.Int
}

func (n *NodeDetails) NodeStatusString() (string, error) {

	switch n.NodeStatus.Int64() {
	case 0:
		return "NotInList", nil
	case 1:
		return "PendingApproval", nil
	case 2:
		return "Approved", nil
	case 3:
		return "Deactivated", nil
	case 4:
		return "Denylisted", nil
	case 5:
		return "Recovery initiated for Denylisted Node", nil
	default:
		return "", errors.New("status not valid")
	}
}
