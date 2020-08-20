package partition

import (
	"encoding/json"
	"fmt"
	"github.com/chengyu-l/ecnode_checker/pkg/chubaofs/proto"
	"github.com/chengyu-l/ecnode_checker/pkg/chubaofs/storage"
	"github.com/chengyu-l/ecnode_checker/pkg/utils"
)

// nodeAddr must be a EcNode
func RequestEcExtents(nodeAddr string, partitionID uint64) (extents []*storage.ExtentInfo, err error) {
	request := utils.NewRequest(proto.OpGetAllWatermarks)
	request.PartitionID = partitionID
	err = utils.DoRequest(request, nodeAddr, 30)
	if err != nil || request.ResultCode != proto.OpOk {
		return nil, fmt.Errorf("request extents fail. resultCode:%v err:%v", request.ResultCode, err)
	}

	err = json.Unmarshal(request.Data, extents)
	if err != nil {
		return nil, fmt.Errorf("unmarshal extents fail. data(%v) err:%v", string(request.Data), err)
	}

	return
}
