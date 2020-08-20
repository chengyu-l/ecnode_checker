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

func (v *Validator) StartValidate() error {
	ep, err := partition.RequestPartition(v.ctx, v.partitionID)
	if err != nil || ep == nil || ep.PartitionID == 0{
		return fmt.Errorf("request partition fail. err:%v", err)
	}

	ecNodeAddr := ep.Hosts[0]
	extents, err := partition.RequestEcExtents(ecNodeAddr, v.partitionID)
	if err != nil || extents == nil {
		return fmt.Errorf("request extens fail. err:%v", err)
	}

	logger := logger.NewLogger(fmt.Sprintf("ep_%v", v.partitionID))
	for _, extent := range extents {
		err = v.doValidate(ecNodeAddr, extent)
		if err != nil {
			// validate fail
			msg := fmt.Sprintf("ep(%v) node(%v) extent(%v) size(%v) validate fail",
				v.partitionID, ecNodeAddr, extent.FileID, extent.Size)
			logger.Record(msg)
		} else {
			// validate success
			msg := fmt.Sprintf("ep(%v) node(%v) extent(%v) size(%v) validate success",
				v.partitionID, ecNodeAddr, extent.FileID, extent.Size)
			logger.Record(msg)
		}
	}

	logger.Record("validate finished")
	logger.Close()
	return nil
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
