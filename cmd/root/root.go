package root

import (
	"github.com/chengyu-l/ecnode_checker/pkg/ecnode"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "ecnode",
}

var Context *utils.Context

func init() {
	Context = utils.NewContext()
}
