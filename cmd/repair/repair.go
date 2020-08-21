package repair

import (
	"fmt"
	"github.com/chengyu-l/ecnode_checker/cmd/root"
	"github.com/chengyu-l/ecnode_checker/pkg/repair"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var Cmd = &cobra.Command{
	Use:   "repair",
	Short: "repair EcExtent of EcPartition on EcNode",
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
	command.Flags().StringVar(&cfg.extentId, "extentId", "", "If set extentId, it will only repair this extent, otherwise, repair all extents in this EcPartition")
}

func init() {
	addCheckerConfigFlags(Cmd)
}

func startValidate(cmd *cobra.Command, args []string) error {
	partitionID, _ := strconv.ParseUint(cfg.partitionId, 10, 0)
	extentID, _ := strconv.ParseUint(cfg.extentId, 10, 0)
	newRepair, err := repair.NewRepair(root.Context, partitionID, extentID)
	if err != nil {
		fmt.Printf("NewValidator err:%v\n", err)
		os.Exit(1)
	}

	isSuccess, err := newRepair.StartRepair()
	if err != nil || !isSuccess {
		fmt.Printf("repair fail. err:%v\n", err)
		os.Exit(1)
	}

	fmt.Println("repair success")
	return nil
}
