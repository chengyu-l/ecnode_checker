package validator

import (
	"fmt"
	"github.com/chengyu-l/ecnode_checker/cmd/root"
	"github.com/chengyu-l/ecnode_checker/pkg/validator"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var Cmd = &cobra.Command{
	Use:   "validate",
	Short: "validate EcExtent of EcPartition on EcNode",
	RunE:  startValidate,
}

var (
	cfg = &config{}
)

type config struct {
	partitionId string
	extentId    string
}

func addCheckerConfigFlags(command *cobra.Command) {
	command.Flags().StringVar(&cfg.partitionId, "partitionId", "", "partitionId")
	command.MarkFlagRequired("partitionId")
	command.Flags().StringVar(&cfg.extentId, "extentId", "", "extentId")
	command.MarkFlagRequired("extentId")
}

func init() {
	addCheckerConfigFlags(Cmd)
}

func startValidate(cmd *cobra.Command, args []string) error {
	partitionID, _ := strconv.ParseUint(cfg.partitionId, 10, 0)
	extentID, _ := strconv.ParseUint(cfg.extentId, 10, 0)
	newValidator, err := validator.NewValidator(root.Context, partitionID, extentID)
	if err != nil {
		fmt.Printf("NewValidator err:%v\n", err)
		os.Exit(1)
	}

	err = newValidator.StartValidate()
	if err != nil {
		fmt.Printf("StartValidate err:%v\n", err)
		os.Exit(1)
	}

	fmt.Println("finished")
	return nil
}
