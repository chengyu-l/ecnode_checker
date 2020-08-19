package main

import (
	"fmt"
	"github.com/chengyu-l/ecnode_checker/cmd/changemember"
	"github.com/chengyu-l/ecnode_checker/cmd/checker"
	"github.com/chengyu-l/ecnode_checker/cmd/root"
	"github.com/chengyu-l/ecnode_checker/cmd/validator"
)

func main() {
	addCommands()
	if err := root.Cmd.Execute(); err != nil {
		fmt.Printf("checker error: %+v\n", err)
	}
}

func addCommands() {
	root.Cmd.AddCommand(
		checker.Cmd,
		validator.Cmd,
		changemember.Cmd,
	)
}
