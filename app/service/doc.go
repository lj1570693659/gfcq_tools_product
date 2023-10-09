package service

import (
	"context"
	"errors"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/lj1570693659/gfcq_tools_product/app/model"
	"github.com/lj1570693659/gfcq_tools_product/boot"
	"gopkg.in/mgo.v2/bson"
)

// Doc 职级信息管理服务
var Doc = docService{}

type docService struct{}

func (s *docService) GetList(ctx context.Context, input *model.DocListInput) ([]model.ObjCateList, error) {
	pro := make([]model.ObjCateList, 0)
	selector := bson.M{}
	if len(input.CNAME) > 0 {
		selector = bson.M{"cname": bson.M{"$regex": bson.RegEx{Pattern: input.CNAME, Options: "im"}}}
	}
	err := boot.DocMongoDB.C("product_cate").Find(selector).Sort("cname").All(&pro)

	if len(pro) > 0 {
		for _, cate := range pro {
			if len(cate.Children) > 0 {
				for _, child := range cate.Children {
					compute := model.NormDiffInfo{}
					err := boot.DocMongoDB.C("product_compute").Find(bson.M{"cid": child.CID}).One(&compute)
					if err == nil {
						child.ShouldSubmitNumber = compute.ShouldSubmitNumber
						child.RealSubmitNumber = compute.RealSubmitNumber
					}
				}
			}
		}
	}
	return pro, err

}

// GetTreeTask TODO
func (s *docService) GetTreeTask(ctx context.Context, input *model.DoctListInput) ([]*model.ProductPlan, error) {
	pro := model.ObjProduct{}
	selector := bson.M{}
	if len(input.CNAME) > 0 {
		selector = bson.M{"cname": bson.M{"$regex": bson.RegEx{Pattern: input.CNAME, Options: "im"}}}
	}
	err := boot.DocMongoDB.C("product_list").Find(selector).Sort("cname").One(&pro)
	return pro.Plan, err

}

func (s *docService) GetComputeByPro(ctx context.Context, input *model.DocComputeInput) ([]model.NormDiffInfo, error) {
	computeResult := make([]model.NormDiffInfo, 0)
	compute := model.NormDiffInfo{}
	selector := bson.M{"cid": input.CID}
	if len(input.CNAME) > 0 {
		selector = bson.M{"cname": bson.M{"$regex": bson.RegEx{Pattern: input.CNAME, Options: "im"}}}
	}
	// 查询对比后数据信息
	err := boot.DocMongoDB.C("product_compute").Find(selector).One(&compute)
	if err != nil {
		if err.Error() == "not found" {
			return computeResult, errors.New("当前项目未同步PLM数据，请先完善交付物上传")
		}
		return computeResult, err
	}
	// 查询项目信息
	product := model.ObjProduct{}
	err = boot.DocMongoDB.C("product_list").Find(selector).One(&product)
	if err != nil {
		return computeResult, err
	}

	// 项目名称 项目层级统计
	computeResult = append(computeResult, model.NormDiffInfo{
		CName:               compute.CName,
		ProjectID:           compute.ProjectID,
		ChangeTime:          compute.ChangeTime,
		ShouldSubmitNumber:  compute.ShouldSubmitNumber,
		ShouldSubmitList:    compute.ShouldSubmitList,
		RealSubmitNumber:    compute.RealSubmitNumber,
		RealSubmitList:      compute.RealSubmitList,
		NormAndRealDiffList: compute.NormAndRealDiffList,
	})

	// 阀点名称 阀点层级统计
	for _, v := range product.StepList {
		pipeline := []bson.M{
			bson.M{"$unwind": bson.M{"path": "$shouldSubmitList"}},
			bson.M{
				"$match": bson.M{
					"cid":                        input.CID,
					"shouldSubmitList.step_name": v,
				},
			},
			bson.M{
				"$project": bson.M{
					"shouldSubmitList": 1,
				},
			},
		}
		// 应该提交数据
		getShouldResult := make([]bson.M, 0)
		boot.DocMongoDB.C("product_compute").Pipe(pipeline).All(&getShouldResult)
		shouldResult := make([]*model.NormFileInfo, 0)
		if len(getShouldResult) > 0 {
			for _, sv := range getShouldResult {
				info := &model.NormFileInfo{}
				gconv.Struct(sv["shouldSubmitList"], &info)
				shouldResult = append(shouldResult, info)
			}
		}
		// 实际提交数据
		getRealInfo := make([]bson.M, 0)
		realPipeline := []bson.M{
			bson.M{"$unwind": bson.M{"path": "$realSubmitList"}},
			bson.M{
				"$match": bson.M{
					"cid":                      input.CID,
					"realSubmitList.step_name": v,
				},
			},
			bson.M{
				"$project": bson.M{
					"realSubmitList": 1,
				},
			},
		}
		boot.DocMongoDB.C("product_compute").Pipe(realPipeline).All(&getRealInfo)
		realResult := make([]*model.MbpProjectTaskFile, 0)
		if len(getRealInfo) > 0 {
			for _, rv := range getRealInfo {
				info := &model.MbpProjectTaskFile{}
				gconv.Struct(rv["realSubmitList"], &info)
				realResult = append(realResult, info)
			}
		}

		// 实际提交数据
		getDiffInfo := make([]bson.M, 0)
		diffPipeline := []bson.M{
			bson.M{"$unwind": bson.M{"path": "$normAndRealDiffInfo"}},
			bson.M{
				"$match": bson.M{
					"cid": input.CID,
					"normAndRealDiffInfo.norm_file_info.step_name": v,
				},
			},
			bson.M{
				"$project": bson.M{
					"normAndRealDiffInfo": 1,
				},
			},
		}
		boot.DocMongoDB.C("product_compute").Pipe(diffPipeline).All(&getDiffInfo)
		diffResult := make([]model.NormAndRealDiffInfo, 0)
		if len(getDiffInfo) > 0 {
			for _, rv := range getDiffInfo {
				info := model.NormAndRealDiffInfo{}
				gconv.Struct(rv["normAndRealDiffInfo"], &info)
				diffResult = append(diffResult, info)
			}
		}

		info := model.NormDiffInfo{
			CName:               v,
			ProjectID:           compute.ProjectID,
			ChangeTime:          compute.ChangeTime,
			ShouldSubmitNumber:  len(shouldResult),
			ShouldSubmitList:    shouldResult,
			RealSubmitNumber:    len(realResult),
			RealSubmitList:      realResult,
			NormAndRealDiffList: diffResult,
		}
		computeResult = append(computeResult, info)
	}

	return computeResult, err

}

func (s *docService) GetStatistics(ctx context.Context, input *model.DocComputeInput) (model.ProductStageLint, error) {
	var result = model.ProductStageLint{
		StageName:   make([]string, 0),
		StageQuota:  make([]float64, 0),
		StageBudget: make([]float64, 0),
	}
	data, err := s.GetComputeByPro(ctx, input)
	if err != nil {
		return result, err
	}
	for k, v := range data {
		if k == 0 {
			continue
		}
		result.StageName = append(result.StageName, v.CName)
		result.StageQuota = append(result.StageQuota, float64(v.RealSubmitNumber))
		result.StageBudget = append(result.StageBudget, float64(v.ShouldSubmitNumber))
	}
	return result, nil
}

func (s *docService) GetProductStatistics(ctx context.Context) (model.ProductStatistics, error) {
	result := model.ProductStatistics{
		ProductName:        make([]string, 0),
		ShouldSubmitNumber: make([]int, 0),
		RealSubmitNumber:   make([]int, 0),
		LackSubmitNumber:   make([]int, 0),
	}

	computeResult := make([]model.NormDiffInfo, 0)

	// 查询对比后数据信息
	err := boot.DocMongoDB.C("product_compute").Find(nil).All(&computeResult)
	if err != nil {
		return result, err
	}
	if len(computeResult) > 0 {
		for _, v := range computeResult {
			result.ProductName = append(result.ProductName, v.CName)
			result.ShouldSubmitNumber = append(result.ShouldSubmitNumber, v.ShouldSubmitNumber)
			result.RealSubmitNumber = append(result.RealSubmitNumber, v.RealSubmitNumber)
			result.LackSubmitNumber = append(result.LackSubmitNumber, v.ShouldSubmitNumber-v.RealSubmitNumber)
		}
	}
	return result, nil
}

func (s *docService) GetDepartStatistics(ctx context.Context, input *model.DocComputeInput) ([]model.DutyPartStatistics, error) {
	diffResult := make([]model.DutyPartStatistics, 0)
	compute := model.NormDiffInfo{}
	selector := bson.M{"cid": input.CID}
	if len(input.CNAME) > 0 {
		selector = bson.M{"cname": bson.M{"$regex": bson.RegEx{Pattern: input.CNAME, Options: "im"}}}
	}
	// 查询对比后数据信息
	err := boot.DocMongoDB.C("product_compute").Find(selector).One(&compute)
	if err != nil {
		if err.Error() == "not found" {
			return diffResult, errors.New("当前项目未同步PLM数据，请先完善交付物上传")
		}
		return diffResult, err
	}

	// 查询项目信息
	product := model.ObjProduct{}
	err = boot.DocMongoDB.C("product_list").Find(selector).One(&product)
	if err != nil {
		return diffResult, err
	}

	// 项目层级统计
	diffCount, err := s.getStatisticsByDepart(bson.M{"$match": bson.M{"cid": input.CID}}, bson.M{})
	if err != nil {
		return diffResult, err
	}
	diffResult = append(diffResult, model.DutyPartStatistics{
		StepName:    product.CNAME,
		CountDetail: diffCount,
	})

	// 阀点名称 阀点层级统计
	for _, v := range product.StepList {
		diffCount, err = s.getStatisticsByDepart(bson.M{
			"$match": bson.M{
				"cid": input.CID,
			},
		}, bson.M{
			"$match": bson.M{
				"realSubmitList.step_name": v,
			},
		})
		if err != nil {
			return diffResult, err
		}
		diffResult = append(diffResult, model.DutyPartStatistics{
			StepName:    v,
			CountDetail: diffCount,
		})
	}

	return diffResult, nil
}

// getStatisticsByDepart 按条件统计各部门提交数量汇总
func (s *docService) getStatisticsByDepart(where1, where2 bson.M) ([]model.DutyPartCount, error) {
	// 实际提交数据
	diffCount := make([]model.DutyPartCount, 0)
	getRealInfo := make([]bson.M, 0)
	realPipeline := []bson.M{
		where1,
		bson.M{"$unwind": bson.M{"path": "$realSubmitList"}},
	}
	if !g.IsEmpty(where2) {
		realPipeline = append(realPipeline, where2)
	}
	realPipeline = append(realPipeline, bson.M{
		"$group": bson.M{"_id": bson.M{"duty_depart": "$realSubmitList.duty_depart"}, "count": bson.M{"$sum": 1}},
	})
	boot.DocMongoDB.C("product_compute").Pipe(realPipeline).All(&getRealInfo)
	if len(getRealInfo) > 0 {
		for _, rv := range getRealInfo {
			diffCount = append(diffCount, model.DutyPartCount{
				Count:      gconv.Int(rv["count"]),
				DutyDepart: gconv.String(gconv.Map(rv["_id"])["duty_depart"]),
			})
		}
	}
	return diffCount, nil
}
