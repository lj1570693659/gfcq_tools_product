package service

import (
	"context"
	"github.com/lj1570693659/gfcq_tools_product/app/model"
	"github.com/lj1570693659/gfcq_tools_product/library/util"
)

// Compute 交付物分析计算类
var Compute = computeService{}

type computeService struct{}

// ComputeAnalysis ComputeAnalysis 分析交付物差距
// norm 标准文件
// userSubmit 用户提交文件
func (s *computeService) ComputeAnalysis(ctx context.Context, norm []*model.NormFileInfo, userSubmit []*model.MbpProjectTaskFile) (model.NormDiffInfo, error) {
	result := model.NormDiffInfo{
		ShouldSubmitNumber: len(norm),
		ShouldSubmitList:   norm,
		RealSubmitNumber:   len(userSubmit),
		RealSubmitList:     userSubmit,
	}
	if len(norm) == 0 || len(userSubmit) == 0 {
		return result, nil
	}

	submitNorm := make([]*model.NormFileInfo, 0)
	submitNormName := make([]string, 0)
	for _, nv := range norm {
		for _, su := range userSubmit {
			if util.GetStrDiff(nv.SplitName, su.CNAME) {
				if !util.CheckIsInArray(submitNormName, nv.Identify) {
					submitNorm = append(submitNorm, nv)
					submitNormName = append(submitNormName, nv.Identify)
				}

			}
		}
	}

	for _, su := range submitNorm {
		if !util.CheckIsInArray(submitNormName, su.Identify) {
			result.LackSubmitList = append(result.LackSubmitList, su)

		}
	}

	return result, nil
}
