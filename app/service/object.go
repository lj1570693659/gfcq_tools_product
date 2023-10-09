package service

import (
	"context"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/lj1570693659/gfcq_tools_product/app/model"
	"github.com/lj1570693659/gfcq_tools_product/boot"
	"github.com/lj1570693659/gfcq_tools_product/library/util"
	"gopkg.in/mgo.v2/bson"
	"sync"
	"time"
)

// Object 职级信息管理服务
var Object = objectService{}

type objectService struct{}

func (s *objectService) SyncObject(ctx context.Context) error {
	proCate := make([]*model.ObjCate, 0)
	boot.OracleDB.Raw(fmt.Sprintf(model.GetCompanyProductCate, boot.CompanyObjId)).Scan(&proCate)

	// 更新项目目录
	if len(proCate) > 0 {
		for _, ca := range proCate {
			err := s.syncObjCate(ca)
			if g.IsNil(err) {
				// 更新项目分类下项目清单
				err = s.syncObjProduct(ca)
				if !g.IsNil(err) {
					return err
				}
			}
		}
	}

	// 更新项目文件清单
	err := s.syncProductFileList()
	return err
}

// syncObjCate 查询公司下面项目管理版块文档目录
func (s *objectService) syncObjCate(cate *model.ObjCate) error {
	info := model.ObjCate{}
	selector := bson.M{"cid": cate.CID}
	boot.DocMongoDB.C("product_cate").Find(selector).One(&info)
	cateInfo := model.ObjCate{
		CID:   cate.CID,
		CNAME: cate.CNAME,
		CTYPE: cate.CTYPE,
	}
	if g.IsEmpty(info.CNAME) {
		boot.DocMongoDB.C("product_cate").Insert(cateInfo)
	} else {
		err := boot.DocMongoDB.C("product_cate").Update(selector, cateInfo)
		if err != nil {
			g.Log("error").Error(err)
			return err
		}
	}
	return nil
}

// syncObjProduct
func (s *objectService) syncObjProduct(cate *model.ObjCate) error {
	proDate := make([]*model.ObjProduct, 0)
	boot.OracleDB.Raw(fmt.Sprintf(model.GetProductListByCate, cate.CID)).Scan(&proDate)

	selector := bson.M{"cid": cate.CID}
	data := bson.M{"$set": bson.M{"children": proDate}}
	err := boot.DocMongoDB.C("product_cate").Update(selector, data)
	if err != nil {
		g.Log("error").Error(err)
		return err
	}

	if len(proDate) > 0 {
		for _, pro := range proDate {
			info := model.ObjProduct{}
			selector := bson.M{"cid": pro.CID}
			boot.DocMongoDB.C("product_list").Find(selector).One(&info)
			if g.IsEmpty(info.CNAME) {
				boot.DocMongoDB.C("product_list").Insert(pro)
			} else {
				data := bson.M{"$set": bson.M{"name": pro.CNAME, "usestate": pro.CUSESTATE}}
				err = boot.DocMongoDB.C("product_list").Update(selector, data)
				if err != nil {
					g.Log("error").Error(err)
				}
			}
		}
	}

	return nil
}

func (s *objectService) syncProductFileList() error {
	proList := make([]model.ObjProduct, 0)
	boot.DocMongoDB.C("product_list").Find(nil).All(&proList)
	if len(proList) > 0 {
		for _, pro := range proList {
			fileDate := make([]*model.ObjProductInfo, 0)
			boot.OracleDB.Raw(fmt.Sprintf(model.GetProductFileListWithoutFilePath, pro.CID)).Scan(&fileDate)

			data := bson.M{"$set": bson.M{"children": fileDate}}
			err := boot.DocMongoDB.C("product_list").Update(bson.M{"cid": pro.CID}, data)
			if err != nil {
				g.Log("error").Error(err)
			}

		}
	}

	return nil
}

// SyncPlan 更新计划
func (s *objectService) SyncPlan(ctx context.Context) error {
	proList := make([]model.ObjProduct, 0)
	boot.DocMongoDB.C("product_list").Find(nil).All(&proList)
	if len(proList) > 0 {
		wg := sync.WaitGroup{}
		for _, pro := range proList {
			wg.Add(1)
			planFileName := g.Config("config.toml").GetString("deliverable.planFileName")
			proPlanFile := model.ObjCate{}
			boot.OracleDB.Raw(fmt.Sprintf(model.GetProductPlanFile, pro.CID, fmt.Sprintf("%s%s%s", "%", planFileName, "%"))).Scan(&proPlanFile)
			var filePath, jsonPath string
			if len(proPlanFile.CID) != 0 {
				fileUrlTemp := g.Config("config.toml").GetString("deliverable.fileUrl")
				url := fmt.Sprintf(fileUrlTemp, proPlanFile.CID, pro.CID)
				fileName := fmt.Sprintf("%s%s", proPlanFile.CNAME, time.Now().Format("20060102"))
				filePath = fmt.Sprintf("./public/%s.mpp", fileName)
				jsonPath = fmt.Sprintf("./public/%s.json", fileName)
				err := util.DownloadFile(url, filePath)
				if err != nil {
					g.Log("error").Error(err)
					return err
				}
			} else {
				fileName := "交付物模板"
				filePath = fmt.Sprintf("./public/%s.mpp", fileName)
				jsonPath = fmt.Sprintf("./public/%s.json", fileName)
			}

			getMppArray, getTreeContent, stepNameList, err := util.GetMppData(filePath, jsonPath)
			if err != nil || len(getMppArray) == 0 {
				g.Log("error").Info(err)
				continue
			}
			err = s.splitNorm(ctx, getMppArray, getTreeContent, stepNameList, pro.CID)
			if err != nil {
				g.Log("error").Error(err)
				return err
			}
			wg.Done()
		}
		wg.Wait()
	}

	return nil
}

// splitNorm
func (s *objectService) splitNorm(ctx context.Context, conArray []*model.ProductPlan, conTree []*model.ProductPlan, stepNameList []string, cid string) error {
	proInfo := model.ObjProduct{}
	selector := bson.M{"cid": cid}
	boot.DocMongoDB.C("product_list").Find(selector).One(&proInfo)
	data := bson.M{"$set": bson.M{"plan": conTree, "plan_list": conArray, "step_list": stepNameList}}
	err := boot.DocMongoDB.C("product_list").Update(selector, data)
	if err != nil {
		g.Log("error").Error(err)
	}
	return err
}

// ComputeDiff 裁剪项目 product_compute
func (s *objectService) ComputeDiff(ctx context.Context) (err error) {
	proList := make([]model.ObjProduct, 0)
	boot.DocMongoDB.C("product_list").Find(nil).All(&proList)
	if len(proList) > 0 {
		wg := sync.WaitGroup{}
		for _, pro := range proList {
			if len(pro.Plan) > 0 {
				wg.Add(1)
				result := model.NormDiffInfo{}
				children := s.syncStepName(pro.Children, pro.StepList)
				result, err = s.computeDiffBySingle(pro.PlanList, children)

				if err == nil {
					s.syncObjProductCompute(ctx, pro, result)
				} else {
					g.Log("error").Info(err)
				}
				wg.Done()
			}
		}
		wg.Wait()
	}
	return err
}
func (s *objectService) syncStepName(fileList []model.ObjProductInfo, stepList []string) []model.ObjProductInfo {
	list := make([]model.ObjProductInfo, 0)
	if len(fileList) == 0 {
		return fileList
	}
	for _, v := range fileList {
		info := v
		info.FXMJD = util.GetStepName(v.FXMJD, stepList)
		list = append(list, info)
	}
	return list
}
func (s *objectService) computeDiffBySingle(plan []*model.ProductPlan, fileList []model.ObjProductInfo) (result model.NormDiffInfo, err error) {
	shouldSubmitList := make([]*model.NormFileInfo, 0)
	normAndRealDiffInfo := make([]model.NormAndRealDiffInfo, 0)
	realSubmitList := make([]*model.MbpProjectTaskFile, 0)
	submitNorm := make([]*model.NormFileInfo, 0)
	submitNormSymbol := make([]string, 0)
	result = model.NormDiffInfo{
		ShouldSubmitNumber:  len(shouldSubmitList),
		ShouldSubmitList:    shouldSubmitList,
		RealSubmitNumber:    len(realSubmitList),
		RealSubmitList:      realSubmitList,
		NormAndRealDiffList: normAndRealDiffInfo,
	}
	if len(plan) == 0 {
		g.Log("error").Info("上传交付物为空")
		return result, nil
	}

	for _, pl := range plan {
		if pl.Cid > 0 || pl.Active == 0 {
			continue
		}

		mbpProjectTaskFile := &model.MbpProjectTaskFile{}
		if isUse, fileInfo := s.fileIsUseNorm(pl, fileList); isUse {
			if !util.CheckIsInArray(submitNormSymbol, fileInfo.CSYMBOL) {
				normFileInfo := &model.NormFileInfo{
					Name: pl.Name,
				}
				submitNorm = append(submitNorm, normFileInfo)
				submitNormSymbol = append(submitNormSymbol, fileInfo.CSYMBOL)
				mbpProjectTaskFile = &model.MbpProjectTaskFile{
					CNAME:      fileInfo.CNAME,
					CVERSION:   fileInfo.CVERSION,
					CSYMBOL:    fileInfo.CSYMBOL,
					StepName:   fileInfo.FXMJD,
					DutyDepart: fileInfo.FGKBM,
				}
				realSubmitList = append(realSubmitList, mbpProjectTaskFile)
			}
		}

		shouldSubmit := &model.NormFileInfo{
			Name:     pl.Name,
			StepName: pl.StepName, // 阀点
		}
		shouldSubmitList = append(shouldSubmitList, shouldSubmit)
		diffInfo := model.NormAndRealDiffInfo{
			ShouldSubmit: shouldSubmit,
			RealSubmit:   mbpProjectTaskFile,
		}
		if len(diffInfo.RealSubmit.CNAME) > 0 {
			diffInfo.IsSubmit = true
		}
		normAndRealDiffInfo = append(normAndRealDiffInfo, diffInfo)
	}

	result = model.NormDiffInfo{
		ShouldSubmitNumber:  len(shouldSubmitList),
		ShouldSubmitList:    shouldSubmitList,
		RealSubmitNumber:    len(realSubmitList),
		RealSubmitList:      realSubmitList,
		NormAndRealDiffList: normAndRealDiffInfo,
	}

	return result, err
}

func (s *objectService) fileIsUseNorm(plan *model.ProductPlan, fileList []model.ObjProductInfo) (check bool, info model.ObjProductInfo) {
	for _, su := range fileList {
		if su.FXMJD == plan.StepName {
			slit := util.GetSplitStr(plan.Name)
			check = util.GetStrDiff(slit, su.CNAME)
			if check {
				g.Log("doc").Info("关键词匹配,标准分析：%v，实际提交关键词：%s", slit, su.CNAME)
				return check, su
			}
		}
	}
	return false, model.ObjProductInfo{}
}

func (s *objectService) syncObjProductCompute(ctx context.Context, product model.ObjProduct, computeResult model.NormDiffInfo) (err error) {
	selector := bson.M{"cid": product.CID}
	info := model.NormDiffInfo{}
	boot.DocMongoDB.C("product_compute").Find(selector).One(&info)
	data := model.NormDiffInfo{
		CID:                 product.CID,
		CName:               product.CNAME,
		ProjectID:           product.CID,
		ChangeTime:          time.Now().Format("2006-01-02 15:04:05"),
		ShouldSubmitNumber:  computeResult.ShouldSubmitNumber,
		RealSubmitNumber:    computeResult.RealSubmitNumber,
		ShouldSubmitList:    computeResult.ShouldSubmitList,
		RealSubmitList:      computeResult.RealSubmitList,
		LackSubmitList:      computeResult.LackSubmitList,
		NormAndRealDiffList: computeResult.NormAndRealDiffList,
	}
	if g.IsEmpty(info.CID) {
		err = boot.DocMongoDB.C("product_compute").Insert(data)
	} else {
		err = boot.DocMongoDB.C("product_compute").Update(selector, data)
	}
	if err != nil {
		g.Log("error").Error(err)
	}
	return err
}
