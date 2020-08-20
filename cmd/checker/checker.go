package checker

import (
	"fmt"
	"github.com/chengyu-l/ecnode_checker/cmd/root"
	"github.com/chengyu-l/ecnode_checker/pkg/checker"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
)

var Cmd = &cobra.Command{
	Use:  "check",
	Short: "check EcExtent of EcPartition on EcNode",
	RunE: startCheck,
}

var (
	cfg = &config{}
)

type config struct {
	partitionId string
	hosts       string
}

func addCheckerConfigFlags(command *cobra.Command) {
	command.Flags().StringVar(&cfg.partitionId, "partitionId", "", "partitionId")
	command.MarkFlagRequired("partitionId")
	command.Flags().StringVar(&cfg.hosts, "hosts", "", "EcPartition hosts. eg: 1.1.1.1:1001,1.1.1.2:1002,1.1.1.3:1003,1.1.1.4:1004,1.1.1.5:1005,1.1.1.6:1006")
	command.MarkFlagRequired("hosts")
}

func init() {
	addCheckerConfigFlags(Cmd)
}

func startCheck(cmd *cobra.Command, args []string) error {
	partitionID, _ := strconv.ParseUint(cfg.partitionId, 10, 0)
	checker, err := checker.NewChecker(root.Context, partitionID, strings.Split(cfg.hosts, ","))
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
