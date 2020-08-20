package checker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	CONFIG_FILE = "ecnode.json"
)

type Context struct {
	MasterAddr []string `json:"masterAddr"`
}

func NewContext() *Context {
	cfgFile, err := os.Open(CONFIG_FILE)
	if err != nil {
		fmt.Printf("ecnode config file is not exist. error:%v\n", err)
		os.Exit(1)
	}

	cfgFileBytes, err := ioutil.ReadAll(cfgFile)
	if err != nil {
		fmt.Printf("fail to open ecnode config file. error:%v\n", err)
		os.Exit(1)
	}

	cxt := &Context{}
	err = json.Unmarshal(cfgFileBytes, cxt)
	if err != nil {
		fmt.Printf("fail to unmarshal to Context. error:%v\n", err)
		os.Exit(1)
	}

	return cxt
}
