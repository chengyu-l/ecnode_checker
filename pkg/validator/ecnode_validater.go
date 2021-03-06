package validator

import (
	"fmt"
	"github.com/chengyu-l/ecnode_checker/pkg/checker"
	"github.com/chengyu-l/ecnode_checker/pkg/chubaofs/proto"
	"github.com/chengyu-l/ecnode_checker/pkg/chubaofs/storage"
	"github.com/chengyu-l/ecnode_checker/pkg/logger"
	"github.com/chengyu-l/ecnode_checker/pkg/partition"
	"github.com/chengyu-l/ecnode_checker/pkg/utils"
)

type Validator struct {
	partitionID uint64
	ctx         *checker.Context
}

func NewValidator(ctx *checker.Context, partitionID uint64) (*Validator, error) {
	return &Validator{
		partitionID: partitionID,
		ctx:         ctx,
	}, nil
}

func (v *Validator) StartValidate() (bool, error) {
	ep, err := partition.RequestEcPartition(v.ctx, v.partitionID)
	if err != nil || ep == nil || ep.PartitionID == 0 {
		return false, fmt.Errorf("request partition fail. err:%v", err)
	}

	ecNodeAddr := ep.Hosts[0]
	extents, err := partition.RequestEcExtents(ecNodeAddr, v.partitionID)
	if err != nil || extents == nil {
		return false, fmt.Errorf("request extens fail. err:%v", err)
	}

	l := logger.NewLogger(fmt.Sprintf("validate_ep_%v", v.partitionID))
	isSuccess := false
	for _, extent := range extents {
		err = v.doValidate(ecNodeAddr, extent)
		if err != nil {
			// validate fail
			msg := fmt.Sprintf("ep(%v) node(%v) extent(%v) size(%v) validate fail\n",
				v.partitionID, ecNodeAddr, extent.FileID, extent.Size)
			l.Record(msg)
		} else {
			// validate success
			msg := fmt.Sprintf("ep(%v) node(%v) extent(%v) size(%v) validate success\n",
				v.partitionID, ecNodeAddr, extent.FileID, extent.Size)
			l.Record(msg)
			isSuccess = true
		}
	}

	l.Close()
	return isSuccess, nil
}

func (v *Validator) doValidate(ecNodeAddr string, extent *storage.ExtentInfo) error {
	request := utils.NewRequest(proto.OpValidateEcDataPartition)
	request.PartitionID = v.partitionID
	request.ExtentID = extent.FileID
	err := utils.DoRequest(request, ecNodeAddr, 500)
	if err != nil || request.ResultCode != proto.OpOk {
		return fmt.Errorf("request validate ecpartition fail. resultCode:%v err:%v", request.ResultCode, err)
	}

	return nil
}
