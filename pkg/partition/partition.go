package partition

import (
	"encoding/json"
	"fmt"
	"github.com/chengyu-l/ecnode_checker/pkg/checker"
	"github.com/chengyu-l/ecnode_checker/pkg/chubaofs/proto"
	"github.com/chengyu-l/ecnode_checker/pkg/utils"
)

func RequestEcPartition(ctx *checker.Context, partitionID uint64) (ep *proto.EcDataPartitionInfo, err error) {
	for i := 0; i < len(ctx.MasterAddr); i++ {
		url := fmt.Sprintf("http://%v/ecPartition/get?id=%v", ctx.MasterAddr[i], partitionID)
		data, err := utils.HttpGetRequest(url)
		if err != nil {
			continue
		}

		fmt.Printf("%s\n", data)
		ep = &proto.EcDataPartitionInfo{}
		reply := &proto.HTTPReply{}
		reply.Data = ep
		err = json.Unmarshal(data, reply)
		if err != nil {
			continue
		}

		if reply.Code != 0 {
			err = fmt.Errorf("reply code not ok, response:%v", string(data))
			continue
		}

		if ep.PartitionID != 0 {
			break
		}
	}

	return ep, err
}
