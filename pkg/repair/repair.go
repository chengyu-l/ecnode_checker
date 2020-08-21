package repair

import (
	"fmt"
	"github.com/chengyu-l/ecnode_checker/pkg/checker"
	"github.com/chengyu-l/ecnode_checker/pkg/chubaofs/proto"
	"github.com/chengyu-l/ecnode_checker/pkg/chubaofs/storage"
	"github.com/chengyu-l/ecnode_checker/pkg/logger"
	"github.com/chengyu-l/ecnode_checker/pkg/partition"
	"github.com/chengyu-l/ecnode_checker/pkg/utils"
)

type Repair struct {
	partitionID uint64
	extentID    uint64
	ctx         *checker.Context
}

func NewRepair(ctx *checker.Context, partitionID, extentID uint64) (*Repair, error) {
	return &Repair{
		partitionID: partitionID,
		extentID:    extentID,
		ctx:         ctx,
	}, nil
}

func (v *Repair) StartRepair() (bool, error) {
	ep, err := partition.RequestEcPartition(v.ctx, v.partitionID)
	if err != nil || ep == nil || ep.PartitionID == 0 {
		return false, fmt.Errorf("request partition fail. err:%v", err)
	}

	ecNodeAddr := ep.Hosts[0]
	extents, err := partition.RequestEcExtents(ecNodeAddr, v.partitionID)
	if err != nil || extents == nil {
		return false, fmt.Errorf("request extens fail. err:%v", err)
	}

	l := logger.NewLogger(fmt.Sprintf("repair_ep_%v", v.partitionID))
	isSuccess := false
	for _, extent := range extents {
		if v.extentID != 0 && v.extentID != extent.FileID {
			continue
		}

		err = v.doRepair(ecNodeAddr, extent)
		if err != nil {
			// validate fail
			msg := fmt.Sprintf("ep(%v) node(%v) extent(%v) size(%v) repair fail\n",
				v.partitionID, ecNodeAddr, extent.FileID, extent.Size)
			l.Record(msg)
		} else {
			// validate success
			msg := fmt.Sprintf("ep(%v) node(%v) extent(%v) size(%v) repair success\n",
				v.partitionID, ecNodeAddr, extent.FileID, extent.Size)
			l.Record(msg)
			isSuccess = true
		}
	}

	l.Close()
	return isSuccess, nil
}

func (v *Repair) doRepair(ecNodeAddr string, extent *storage.ExtentInfo) error {
	request := utils.NewRequest(proto.OpEcExtentRepair)
	request.PartitionID = v.partitionID
	request.ExtentID = extent.FileID
	err := utils.DoRequest(request, ecNodeAddr, 500)
	if err != nil || request.ResultCode != proto.OpOk {
		return fmt.Errorf("request validate ecpartition fail. resultCode:%v err:%v", request.ResultCode, err)
	}

	return nil
}
