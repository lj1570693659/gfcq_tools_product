package service

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/lj1570693659/gfcq_product_kpi/app/model"
	"github.com/lj1570693659/gfcq_tools_product/consts"
	"github.com/lj1570693659/gfcq_tools_product/library/response"
	"net/http"
)

// Middleware 中间件管理服务
var Middleware = middlewareService{}

type middlewareService struct{}

// Ctx 自定义上下文对象
func (s *middlewareService) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Session: r.Session,
	}
	Context.Init(r, customCtx)
	ctx := r.Context()
	if user := Session.GetUser(ctx); user == nil || len(user.WorkNumber) > 0 {
		// 未登录
		token := r.Cookie.Get(consts.TokenName)
		if len(token) == 0 {
			response.JsonExit(r, http.StatusForbidden, "请先登录")
		}
		Session.SetUser(ctx, &model.User{
			WorkNumber: token,
		})
	}

	// 执行下一步请求逻辑
	r.Middleware.Next()
}

// LoggedIn 鉴权中间件，验证是否登录
func (s *middlewareService) LoggedIn(r *ghttp.Request) {
	r.Middleware.Next()
	//if User.IsSignedIn(r.Context()) {
	//	r.Middleware.Next()
	//} else {
	//	response.JsonExit(r, http.StatusForbidden, "")
	//}
}

// Role 鉴权中间件，验证是否在允许角色组内
//func (s *middlewareService) Role(r *ghttp.Request) {
//	ok, err := Casbin.CheckAuth(r.Context(), Context.Get(r.Context()).User, r, ROLE)
//	if err != nil {
//		response.JsonExit(r, http.StatusForbidden, err.Error())
//	}
//	if ok {
//		r.Middleware.Next()
//	} else {
//		response.JsonExit(r, http.StatusForbidden, "当前用户权限不足")
//	}
//}
//
//// BusinessRole 鉴权中间件，验证是否在项目组内
//func (s *middlewareService) BusinessRole(r *ghttp.Request) {
//	ok, err := Casbin.CheckAuth(r.Context(), Context.Get(r.Context()).User, r, BUSINESS_ROLE)
//	if err != nil {
//		response.JsonExit(r, http.StatusForbidden, err.Error())
//	}
//	if ok {
//		r.Middleware.Next()
//	} else {
//		response.JsonExit(r, http.StatusForbidden, "当前用户权限不足")
//	}
//}

// CORS 允许接口跨域请求
func (s *middlewareService) CORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = []string{"localhost:8199", "10.24.12.84:8199", "192.168.137.1:8199",
		"10.80.42.65:8199", "localhost:9528", "10.80.42.65:9528", "127.0.0.1:8197", "10.80.28.218:9530", "10.80.28.218:8130"}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}

// SaveUserLog 保存用户操作记录
//func (s *middlewareService) SaveUserLog(r *ghttp.Request) {
//	err := UserLog.SaveLogData(r.Context(), r.Cookie, r.GetRequestMap(), r.Method, r.RequestURI)
//	if err != nil {
//		response.JsonExit(r, response.CreateFailLog, err.Error())
//	}
//
//	r.Middleware.Next()
//}
