package ecmonkey

import (
	"fmt"
	"github.com/chengyu-l/ecnode_checker/pkg/utils"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strconv"
)

var Cmd = &cobra.Command{
	Use:   "ecmonkey",
	Short: "deliberately damage a Extent file on EcNode",
	RunE:  startValidate,
}

var (
	cfg = &config{}
)

type config struct {
	file   string
	offset string
	size   string
}

func addCheckerConfigFlags(command *cobra.Command) {
	command.Flags().StringVar(&cfg.file, "file", "", "extent file path")
	command.MarkFlagRequired("file")
	command.Flags().StringVar(&cfg.offset, "offset", "10", "damage the extent file from the offset")
	command.Flags().StringVar(&cfg.size, "size", "1", "how many data size are damaged in the extent file")
}

func init() {
	addCheckerConfigFlags(Cmd)
}

func startValidate(cmd *cobra.Command, args []string) error {
	exists, err := utils.PathExists(cfg.file)
	if err != nil || !exists {
		fmt.Println(fmt.Sprintf("file(%v) not found, err:%v", cfg.file, err))
		os.Exit(1)
	}

	file, err := os.OpenFile(cfg.file, os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(fmt.Sprintf("open file(%v) fail, err:%v", cfg.file, err))
		os.Exit(1)
	}

	offset, _ := strconv.ParseInt(cfg.offset, 10, 0)
	size, _ := strconv.ParseUint(cfg.size, 10, 0)
	data := make([]byte, size)
	_, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		fmt.Println(fmt.Sprintf("seek file(%v) fail, offset(%v) err:%v", cfg.file, offset, err))
		os.Exit(1)
	}

	_, err = file.Write(data)
	if err != nil {
		fmt.Println(fmt.Sprintf("write file(%v) fail, err:%v", cfg.file, err))
		os.Exit(1)
	}

	_ = file.Sync()
	_ = file.Close()
	fmt.Println(fmt.Sprintf("very happy as a monkey ^_^ . offset(%v) size(%v)", offset, size))
	return nil
}
