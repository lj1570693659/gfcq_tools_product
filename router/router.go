package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/lj1570693659/gfcq_tools_product/app/api"
	"github.com/lj1570693659/gfcq_tools_product/app/service"
)

func init() {
	s := g.Server()
	// 分组路由注册方式
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware.Ctx,
			service.Middleware.CORS,
		)
		group.Group("/timer", func(g *ghttp.RouterGroup) {
			// 定时任务相关
			g.GET("/start", api.Timer.SyncProduct)
		})
		// 系统管理
		group.Group("/plm", func(sg *ghttp.RouterGroup) {
			// 文档管理 TODO
			sg.Group("/account", func(sga *ghttp.RouterGroup) {
				// 个人资料、修改密码、日志 TODO
				// 登录账号相关
				//sga.ALL("/user", api.Product)
			})

			// 项目管理 TODO
			sg.Group("/product", func(pga *ghttp.RouterGroup) {
				// 我管理的项目
				pga.GET("/lists", api.Product.GetList)
				pga.GET("/task", api.Product.GetTreeTask)
				pga.GET("/taskCount", api.Product.GetTreeTaskCount)
				// 维护人员使用接口
				// sync_product 更新项目
				pga.GET("/sync_product", api.Product.SyncProduct)
				// sync_task 更新项目任务
				pga.GET("/sync_task", api.Product.SyncTask)
				// sync_task 更新项目任务下交付物
				pga.GET("/sync_task_doc", api.Product.SyncTaskDoc)
				// sync_task 交付物匹配信息
				pga.GET("/sync_doc_count", api.Product.SyncDocCount)
				// sync_doc_diff 交付物统计信息
				pga.GET("/sync_doc_diff", api.Product.ComputeDiff)
				// fenci_norm 模板分词
				pga.GET("/fenci_norm", api.Product.Norm)
			})
			sg.Group("/doc", func(pga *ghttp.RouterGroup) {
				// 我管理的项目
				pga.GET("/lists", api.Doc.GetList)
				pga.GET("/task", api.Doc.GetTreeTask)
				pga.GET("/compute", api.Doc.GetComputeByPro)
				pga.GET("/statistics", api.Doc.GetStatistics)
				pga.GET("/product_statistics", api.Doc.GetProductStatistics)
				pga.GET("/depart_statistics", api.Doc.GetDepartStatistics)

				// 维护人员使用接口
				// sync_object 更新项目
				pga.GET("/sync_object", api.Object.SyncObject)
				// sync_plan 更新项目计划
				pga.GET("/sync_plan", api.Object.SyncPlan)
				// sync_compute 交付物匹配信息
				pga.GET("/sync_compute", api.Object.ComputeDiff)
			})
		})
	})
}
