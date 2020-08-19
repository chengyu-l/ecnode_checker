package validator

import (
	"fmt"
	"github.com/chengyu-l/ecnode_checker/pkg/chubaofs/proto"
	"github.com/chengyu-l/ecnode_checker/pkg/ecnode"
)

type Validator struct {
	partitionID uint64
	extentID    uint64
	ctx         *ecnode.Context
}

func NewValidator(ctx *ecnode.Context, partitionID, extentID uint64) (*Validator, error) {
	return &Validator{
		partitionID: partitionID,
		extentID:    extentID,
		ctx:         ctx,
	}, nil
}

func (v *Validator) StartValidate() error {
	request := ecnode.NewRequest(proto.OpValidateEcDataPartition)
	request.PartitionID = v.partitionID
	request.ExtentID = v.extentID
	err := ecnode.DoRequest(request, "", 30)
	if err != nil {
		return err
	}

	if request.ResultCode != proto.OpOk {
		return fmt.Errorf("response not ok. resultCode:%v", request.ResultCode)
	}

	return nil
}
