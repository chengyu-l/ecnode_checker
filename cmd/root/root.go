package root

import (
	"github.com/chengyu-l/ecnode_checker/pkg/checker"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "ecnode_checker",
}

var Context *checker.Context

func init() {
	Context = checker.NewContext()
}
