package main

import (
	"fmt"
	"github.com/chengyu-l/cfs-ecnode-checker/pkg"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
)

var CheckerCmd = &cobra.Command{
	Use:  "checker",
	RunE: startCheck,
}

var (
	cfg = &config{}
)

type config struct {
	partitionID string
	hosts       string
}

func addOSDConfigFlags(command *cobra.Command) {
	command.Flags().StringVar(&cfg.partitionID, "partitionId", "", "partitionId")
	command.MarkFlagRequired("partitionId")
	command.Flags().StringVar(&cfg.hosts, "hosts", "", "EcPartition hosts. eg: 1.1.1.1:1001,1.1.1.2:1002,1.1.1.3:1003,1.1.1.4:1004,1.1.1.5:1005,1.1.1.6:1006")
	command.MarkFlagRequired("hosts")
}

func init() {
	addOSDConfigFlags(CheckerCmd)
}

func main() {
	if err := CheckerCmd.Execute(); err != nil {
		fmt.Printf("checker error: %+v\n", err)
	}
}

func startCheck(cmd *cobra.Command, args []string) error {
	partitionID, _ := strconv.Atoi(cfg.partitionID)
	checker, err := pkg.NewChecker(uint64(partitionID), strings.Split(cfg.hosts, ","))
	if err != nil {
		fmt.Printf("NewChecker err:%v\n", err)
		os.Exit(1)
	}

	err = checker.StartCheck()
	if err != nil {
		fmt.Printf("StartCheck err:%v\n", err)
		os.Exit(1)
	}

	fmt.Println("finished")
	return nil
}
