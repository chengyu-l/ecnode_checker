package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/chengyu-l/chubaofs/proto"
	"github.com/chengyu-l/chubaofs/repl"
	"hash/crc32"
	"net"
)

type Checker struct {
	node        string
	partitionID uint64
	hosts       []string
}

func NewChecker(partitionID uint64, hosts []string) (*Checker, error) {
	if len(hosts) != 6 {
		return nil, fmt.Errorf("a partition must have 6 EcNode")
	}

	return &Checker{
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

	request := repl.NewPacket()
	request.PartitionID = c.partitionID
	request.Opcode = proto.OpChangeEcPartitionMembers
	request.ReqID = proto.GenerateRequestID()
	request.Data = marshaled
	request.Size = uint32(len(marshaled))
	request.CRC = crc32.ChecksumIEEE(marshaled)
	err = doRequest(request, c.node, 30)
	if err != nil {
		return
	}

	if request.ResultCode != proto.OpOk {
		return fmt.Errorf("response not ok. resultCode:%v", request.ResultCode)
	}

	return nil
}

func doRequest(request *repl.Packet, addr string, timeoutSec int) (err error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			fmt.Println(fmt.Sprintf("close tcp connection fail, host(%v) error(%v)", addr, err))
		}
	}()

	err = request.WriteToConn(conn)
	if err != nil {
		err = fmt.Errorf("write to host(%v) error(%v)", addr, err)
		return
	}

	err = request.ReadFromConn(conn, timeoutSec) // read the response
	if err != nil {
		err = fmt.Errorf("read from host(%v) error(%v)", addr, err)
		return
	}

	return
}
