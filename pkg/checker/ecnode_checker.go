package checker

import (
	"encoding/json"
	"fmt"
	"github.com/chengyu-l/ecnode_checker/pkg/chubaofs/proto"
	"github.com/chengyu-l/ecnode_checker/pkg/utils"
	"hash/crc32"
)

type Checker struct {
	node        string
	partitionID uint64
	hosts       []string
	ctx         *Context
}

func NewChecker(ctx *Context, partitionID uint64, hosts []string) (*Checker, error) {
	if len(hosts) != 6 {
		return nil, fmt.Errorf("a partition must have 6 EcNode")
	}

	return &Checker{
		ctx:         ctx,
		node:        hosts[0],
		partitionID: partitionID,
		hosts:       hosts,
	}, nil
}

func (c *Checker) StartCheck() (err error) {
	changeRequest := &proto.ChangeEcPartitionMembersRequest{}
	changeRequest.PartitionId = c.partitionID
	changeRequest.Hosts = c.hosts
	task := proto.AdminTask{}
	task.PartitionID = c.partitionID
	task.OpCode = proto.OpChangeEcPartitionMembers
	task.ID = "checker-1"
	task.Request = changeRequest

	marshaled, err := json.Marshal(task)
	if err != nil {
		return err
	}

	request := utils.NewRequest(proto.OpChangeEcPartitionMembers)
	request.PartitionID = c.partitionID
	request.Data = marshaled
	request.Size = uint32(len(marshaled))
	request.CRC = crc32.ChecksumIEEE(marshaled)
	err = utils.DoRequest(request, c.node, 30)
	if err != nil {
		return
	}

	if request.ResultCode != proto.OpOk {
		return fmt.Errorf("response not ok. resultCode:%v", request.ResultCode)
	}

	return nil
}
