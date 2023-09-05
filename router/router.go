package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/lj1570693659/gfcq_tools_product/app/service"
)

func init() {
	s := g.Server()
	// 分组路由注册方式
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware.Ctx,
			service.Middleware.CORS,
			//service.Middleware.SaveUserLog,
		)

		// 系统管理
		group.Group("/system", func(sg *ghttp.RouterGroup) {
			// 账号管理
			sg.Group("/account", func(sga *ghttp.RouterGroup) {
				// 个人资料、修改密码、日志 TODO
				// 登录账号相关
				//sga.ALL("/user", api.User)
			})

		})
	})
}
