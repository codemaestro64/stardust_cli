package cli

import (
	"fmt"
)

type CreateCommand struct {
	args        *Args
	subCommands map[string]func() error
}

const (
	createCommand = "Create"
)

func newCreateCommand(args *Args) Command {
	c := &CreateCommand{
		args: args,
	}
	c.subCommands = map[string]func() error{
		"master-wallet": c.createMasterWallet,
		"btc-wallet":    c.createBTCWallet,
		"eth-wallet":    c.createEthWallet,
		"network":       c.createNetwork,
		"node":          c.createNode,
	}

	return c
}

func (c *CreateCommand) createMasterWallet() error {
	fmt.Println("create master wallet")
	return nil
}

func (c *CreateCommand) createBTCWallet() error {
	fmt.Println("create BTC wallet")
	return nil
}

func (c *CreateCommand) createEthWallet() error {
	fmt.Println("create ETH wallet")
	return nil
}

func (c *CreateCommand) createNetwork() error {
	networkTypeNumbder := c.args.pop()
	if networkTypeNumbder == "" {
		fmt.Println("create network")
		return nil
	}

	fmt.Printf("create network type with number: %s", networkTypeNumbder)
	return nil
}

func (c *CreateCommand) createNode() error {
	nodeHash := c.args.pop()
	if nodeHash == "" {
		return fmt.Errorf("%s%s", createCommand, "-Node: invalid node hash")
	}

	fmt.Println("create node")
	return nil
}

func (c *CreateCommand) Execute() error {
	subCommand := c.args.pop()
	if subCommand == "" {
		return fmt.Errorf("%s: %s", createCommand, "sub command not supplied")
	}

	if subCommand, ok := c.subCommands[subCommand]; ok {
		return subCommand()
	}

	return fmt.Errorf("%s: %s", createCommand, "invalid sub command")
}
