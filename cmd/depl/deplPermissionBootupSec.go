/*
Copyright © 2024 Inderjit Singh
ijsingh82@gmail.com
*/
package depl

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	c "github.com/ijsingh82/qV2Cli/bind"
	"github.com/ijsingh82/qV2Cli/cmd"
	"github.com/ijsingh82/qV2Cli/pkg/model"
	"github.com/ijsingh82/qV2Cli/pkg/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

// DepCmd represents the Dep command
var permCmd = &cobra.Command{
	Use:              "perm",
	TraverseChildren: true,
	Short:            "Used to deploy permisionning on a new chain.",
	Long: `
	Used to deploy permisionning on a new chain.

	qV2Cli depl perm [secretKey] [address] [adminOrg] [newAdminRole] [orgAdminRole] [subOrgDepth] [subOrgBreadth]`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(7)(cmd, args); err != nil {
			return err
		}
		_, err := os.Stat(args[0])
		if err != nil {
			return fmt.Errorf("invalid file path : %s", args[0])
		}
		_, err = strconv.Atoi(args[5])
		if err != nil {
			return fmt.Errorf("subOrgDepth not a number: %s", args[5])
		}
		_, err = strconv.Atoi(args[6])
		if err != nil {
			return fmt.Errorf("subOrgBreadth not a number: %s", args[6])
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		keyFile := args[0]
		pass := args[1]
		adminOrg := args[2]
		newAdminRole := args[3]
		orgAdminRole := args[4]
		subOrgDepth, _ := strconv.Atoi(args[5])
		subOrgBreadth, _ := strconv.Atoi(args[6])

		if pass == "" {
			err := deploy(keyFile, "", adminOrg, newAdminRole, orgAdminRole, subOrgDepth, subOrgBreadth)
			if err != nil {
				return err
			}
		} else {
			err := deploy(keyFile, pass, adminOrg, newAdminRole, orgAdminRole, subOrgDepth, subOrgBreadth)
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {

	deplCmd.AddCommand(permCmd)

}

func deploy(key, add, adminOrg, newAdminRole, orgAdminRole string, subOrgDepth, subOrgBreadth int) error {

	client, err := ethclient.Dial(cmd.Conf.Rpc)
	if err != nil {
		return err
	}
	defer client.Close()

	auth1, err := utils.GetNextAuth(add, key, client)
	if err != nil {
		return err
	}
	upgAdd, upgTx, _, err := c.DeployPermUpgr(auth1, client, common.HexToAddress(add))
	if err != nil {
		return err
	}
	_, err = bind.WaitDeployed(context.Background(), client, upgTx)
	if err != nil {
		_, err := utils.GetFailingMessage(client, upgTx.Hash())
		return err
	}
	fmt.Println("The Upgradable address is :", upgAdd)
	utils.PrintToFile("./upgAddTx", []byte(upgTx.String()))

	auth2, err := utils.GetNextAuth(add, key, client)
	if err != nil {
		return err
	}
	orgManAdd, orgTx, _, err := c.DeployOrgManager(auth2, client, upgAdd)
	if err != nil {
		return err
	}
	_, err = bind.WaitDeployed(context.Background(), client, orgTx)
	if err != nil {
		_, err := utils.GetFailingMessage(client, orgTx.Hash())
		return err
	}
	fmt.Println("The orgManAdd address is :", orgManAdd)
	utils.PrintToFile("./orgManAddTx", []byte(orgTx.String()))

	auth3, err := utils.GetNextAuth(add, key, client)
	if err != nil {
		return err
	}
	roleManAdd, roleTx, _, err := c.DeployRoleManager(auth3, client, upgAdd)
	if err != nil {
		return err
	}
	_, err = bind.WaitDeployed(context.Background(), client, roleTx)
	if err != nil {
		_, err := utils.GetFailingMessage(client, roleTx.Hash())
		return err
	}
	fmt.Println("The roleManAdd address is :", roleManAdd)
	utils.PrintToFile("./roleManAddTx", []byte(roleTx.String()))

	auth4, err := utils.GetNextAuth(add, key, client)
	if err != nil {
		return err
	}
	acctManAdd, acctTx, _, err := c.DeployAcctManager(auth4, client, upgAdd)
	if err != nil {
		return err
	}
	_, err = bind.WaitDeployed(context.Background(), client, acctTx)
	if err != nil {
		_, err := utils.GetFailingMessage(client, acctTx.Hash())
		return err
	}
	fmt.Println("The acctManAdd address is :", acctManAdd)
	utils.PrintToFile("./acctManAddTx", []byte(acctTx.String()))

	auth5, err := utils.GetNextAuth(add, key, client)
	if err != nil {
		return err
	}
	voteManAdd, voteTx, _, err := c.DeployVoterManager(auth5, client, upgAdd)
	if err != nil {
		return err
	}
	_, err = bind.WaitDeployed(context.Background(), client, voteTx)
	if err != nil {
		_, err := utils.GetFailingMessage(client, voteTx.Hash())
		return err
	}
	fmt.Println("The acctManAdd address is :", voteManAdd)
	utils.PrintToFile("./voteManAddTx", []byte(voteTx.String()))

	auth6, err := utils.GetNextAuth(add, key, client)
	if err != nil {
		return err
	}
	nodeManAdd, nodeTx, _, err := c.DeployNodeManager(auth6, client, upgAdd)
	if err != nil {
		return err
	}
	_, err = bind.WaitDeployed(context.Background(), client, nodeTx)
	if err != nil {
		_, err := utils.GetFailingMessage(client, nodeTx.Hash())
		return err
	}
	fmt.Println("The nodeManAdd address is :", nodeManAdd)
	utils.PrintToFile("./nodeManAddTx", []byte(nodeTx.String()))

	auth7, err := utils.GetNextAuth(add, key, client)
	if err != nil {
		return err
	}
	interfaceAdd, interfaceTx, _, err := c.DeployPermInterface(auth7, client, upgAdd)
	if err != nil {
		return err
	}
	_, err = bind.WaitDeployed(context.Background(), client, interfaceTx)
	if err != nil {
		_, err := utils.GetFailingMessage(client, interfaceTx.Hash())
		return err
	}
	fmt.Println("The interfaceAdd address is :", interfaceAdd)
	utils.PrintToFile("./interfaceAddTx", []byte(interfaceTx.String()))

	auth8, err := utils.GetNextAuth(add, key, client)
	if err != nil {
		return err
	}
	implAdd, implTx, _, err := c.DeployPermImpl(auth8, client, upgAdd, orgManAdd, roleManAdd, acctManAdd, voteManAdd, nodeManAdd)
	if err != nil {
		return err
	}
	_, err = bind.WaitDeployed(context.Background(), client, implTx)
	if err != nil {
		_, err := utils.GetFailingMessage(client, implTx.Hash())
		return err
	}
	fmt.Println("The implAdd address is :", implAdd)
	utils.PrintToFile("./implAddTx", []byte(implTx.String()))

	auth9, err := utils.GetNextAuth(add, key, client)
	if err != nil {
		return err
	}
	permInst, err := c.NewPermUpgr(upgAdd, client)
	if err != nil {
		return err
	}
	tx, err := permInst.Init(auth9, interfaceAdd, implAdd)
	if err != nil {
		return err
	}
	recipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return err
	}
	if recipt.Status == 0 {
		_, err := utils.GetFailingMessage(client, tx.Hash())
		return err
	}

	g := model.Config{
		PermissionModel:   "v2",
		UpgradableAddress: upgAdd.String(),
		InterfaceAddress:  interfaceAdd.String(),
		ImplAddress:       implAdd.String(),
		NodeMgrAddress:    nodeManAdd.String(),
		AccountMgrAddress: acctManAdd.String(),
		RoleMgrAddress:    roleManAdd.String(),
		VoterMgrAddress:   voteManAdd.String(),
		OrgMgrAddress:     orgManAdd.String(),
		GuardianAddress:   []string{add},
		NewAdminOrg:       adminOrg,
		NewAdminRole:      newAdminRole,
		OrgAdminRole:      orgAdminRole,
		SubOrgDepth:       subOrgDepth,
		SubOrgBreadth:     subOrgBreadth,
	}
	file, err := json.Marshal(g)
	if err != nil {
		log.Fatal("json marshal error")
	}
	utils.PrintToFile("./permission-config.json", file)

	return nil
}
