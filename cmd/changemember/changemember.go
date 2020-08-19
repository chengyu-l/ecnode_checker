package changemember

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:  "cm",
	RunE: startChange,
}

var (
	cfg = &config{}
)

type config struct {
	partitionID string
	hosts       string
}

func addCheckerConfigFlags(command *cobra.Command) {
	command.Flags().StringVar(&cfg.partitionID, "partitionId", "", "partitionId")
	command.MarkFlagRequired("partitionId")
	command.Flags().StringVar(&cfg.hosts, "hosts", "", "EcPartition hosts. eg: 1.1.1.1:1001,1.1.1.2:1002,1.1.1.3:1003,1.1.1.4:1004,1.1.1.5:1005,1.1.1.6:1006")
	command.MarkFlagRequired("hosts")
}

func init() {
	addCheckerConfigFlags(Cmd)
}

func startChange(cmd *cobra.Command, args []string) error {
	//partitionID, _ := strconv.Atoi(cfg.partitionID)
	//checker, err := checker2.NewChecker(uint64(partitionID), strings.Split(cfg.hosts, ","))
	//if err != nil {
	//	fmt.Printf("NewChecker err:%v\n", err)
	//	os.Exit(1)
	//}
	//
	//err = checker.StartCheck()
	//if err != nil {
	//	fmt.Printf("StartCheck err:%v\n", err)
	//	os.Exit(1)
	//}

	fmt.Println("finished")
	return nil
}

