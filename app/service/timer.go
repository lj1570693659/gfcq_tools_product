package service

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/gctx"
)

// Timer 定时任务
var Timer = timerService{}

type timerService struct{}

func (s *timerService) SyncProduct() error {
	var (
		err error
		ctx = gctx.New()
	)

	/***********项目管理板块******************/
	//  更新项目
	_, err = gcron.Add("0 0 01 * * *", func() {
		Plm.SyncProduct(ctx)
	}, "ProductSyncProduct")
	if err != nil {
		g.Log("error").Warning(fmt.Sprint("ProductSyncProduct.err", err.Error()))
	}
	//  更新项目任务
	_, err = gcron.Add("0 10 01 * * *", func() {
		Plm.SyncTask(ctx)
	}, "ProductSyncTask")
	if err != nil {
		g.Log("error").Warning(fmt.Sprint("ProductSyncTask.err", err.Error()))
	}
	//  更新项目任务下交付物
	_, err = gcron.Add("0 20 01 * * *", func() {
		Plm.SyncTaskDoc(ctx)
	}, "ProductSyncTaskDoc")
	if err != nil {
		g.Log("error").Warning(fmt.Sprint("ProductSyncTaskDoc.err", err.Error()))
	}
	//  交付物匹配信息
	_, err = gcron.Add("0 30 01 * * *", func() {
		Plm.SyncDocCount(ctx)
	}, "ProductSyncDocCount")
	if err != nil {
		g.Log("error").Warning(fmt.Sprint("ProductSyncDocCount.err", err.Error()))
	}

	/***********文档管理板块******************/
	//  更新项目
	_, err = gcron.Add("0 40 01 * * *", func() {
		Object.SyncObject(ctx)
	}, "DocSyncObject")
	if err != nil {
		g.Log("error").Warning(fmt.Sprint("DocSyncObject.err", err.Error()))
	}
	//  更新项目计划
	_, err = gcron.Add("0 50 01 * * *", func() {
		Object.SyncPlan(ctx)
	}, "DocSyncPlan")
	if err != nil {
		g.Log("error").Warning(fmt.Sprint("DocSyncPlan.err", err.Error()))
	}
	//  交付物匹配信息
	_, err = gcron.Add("0 60 01 * * *", func() {
		Object.ComputeDiff(ctx)
	}, "DocComputeDiff")
	if err != nil {
		g.Log("error").Warning(fmt.Sprint("DocComputeDiff.err", err.Error()))
	}

	if err != nil {
		panic(err)
	}
	select {}
}
