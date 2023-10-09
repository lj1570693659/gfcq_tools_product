package service

import (
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/lj1570693659/gfcq_tools_product/app/model"
	"github.com/lj1570693659/gfcq_tools_product/boot"
	"github.com/lj1570693659/gfcq_tools_product/library/util"
	"gopkg.in/mgo.v2/bson"
)

// Deliverable 交付物分析计算类
var Deliverable = deliverableService{}

type deliverableService struct{}

func (s *deliverableService) SplitNorm(ctx context.Context, in *model.ProductTaskCompute) error {
	norm := model.NormList{}
	selector := bson.M{"name": in.CNAME}
	boot.ProductMongoDB.C("deliverable").Find(selector).One(&norm)
	for _, no := range norm.Records {
		no.SplitName = util.GetSplitStr(no.Name)
	}
	data := bson.M{"$set": bson.M{"records": norm.Records}}
	err := boot.ProductMongoDB.C("deliverable").Update(selector, data)
	if err != nil {
		g.Log("error").Error(err)
	}
	return err
}
