package service

import (
	"context"
	"fmt"
	"github.com/lj1570693659/gfcq_tools_product/app/model"
	"github.com/lj1570693659/gfcq_tools_product/boot"
	"github.com/lj1570693659/gfcq_tools_product/library/response"
	"gopkg.in/mgo.v2/bson"
)

// Product 职级信息管理服务
var Product = productService{}

type productService struct{}

func (s *productService) GetList(ctx context.Context, input *model.ProductListInput) (response.GetListResponse, error) {
	pro := make([]model.MBPPROJECT, 0)
	selector := bson.M{}
	if len(input.CNAME) > 0 {
		selector = bson.M{"cname": bson.M{"$regex": bson.RegEx{Pattern: input.CNAME, Options: "im"}}}
	}
	err := boot.ProductMongoDB.C("plm").Find(selector).Sort("cname").Skip((input.Page - 1) * input.Size).Limit(input.Size).All(&pro)

	totalSize, err := boot.ProductMongoDB.C("plm").Find(nil).Count()
	return response.GetListResponse{
		Page:      input.Page,
		Size:      input.Size,
		TotalSize: totalSize,
		Data:      pro,
	}, err

}

func (s *productService) GetTreeTask(ctx context.Context, input *model.ProductTaskListInput) (interface{}, error) {
	pro := model.ProductTask{}
	selector := bson.M{"productId": input.CID}
	err := boot.ProductMongoDB.C("task").Find(selector).One(&pro)

	return pro.TaskList, err

}

func (s *productService) GetTreeTaskCount(ctx context.Context, input *model.ProductTaskListInput) (interface{}, error) {
	pro := model.ProductTask{}
	lists := make([]model.ProductTaskDoc, 0)
	selector := bson.M{"productId": input.CID}
	err := boot.ProductMongoDB.C("task").Find(selector).One(&pro)
	if len(pro.TaskList) > 0 {
		for _, v := range pro.TaskList {
			for _, vt := range v.TaskFile {
				vt.CID = fmt.Sprintf("%s/%s", vt.CID, vt.PROJECTID)
			}
			lists = append(lists, model.ProductTaskDoc{
				CID:           v.CID,
				CVERSION:      v.CPARENTTASKID,
				CNAME:         v.CNAME,
				CSYMBOL:       "",
				TaskFileCount: v.TaskFileCount,
				Children:      v.TaskFile,
			})
		}
	}
	return lists, err
}
