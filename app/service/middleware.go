package service

import (
	"github.com/gogf/gf/net/ghttp"
)

// Middleware 中间件管理服务
var Middleware = middlewareService{}

type middlewareService struct{}

// Ctx 自定义上下文对象
func (s *middlewareService) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	//customCtx := &model.Context{
	//	Session: r.Session,
	//}
	//Context.Init(r, customCtx)
	//if user := Session.GetUser(r.Context()); user != nil {
	//	customCtx.User = &model.ContextUser{}
	//	customCtx.User.UserInfo = &model.UserInfo{
	//		Id:         gconv.Uint(user.Id),
	//		WorkNumber: user.WorkNumber,
	//		Password:   user.Password,
	//	}
	//
	//	// 完善上下文员工信息
	//	employeeInfo, err := Employee.GetOne(r.Context(), &model.EmployeeApiGetOneReq{
	//		model.Employee{
	//			WorkNumber: user.WorkNumber,
	//		},
	//	})
	//	if err != nil && rpctypes.ErrorDesc(err) != sql.ErrNoRows.Error() {
	//		response.JsonExit(r, http.StatusForbidden, err.Error())
	//	}
	//	Context.SetUserEmployee(r.Context(), &employeeInfo.EmployeeInfo)
	//	Context.SetUserDepartment(r.Context(), employeeInfo.DepartmentInfo)
	//	Context.SetUserJob(r.Context(), employeeInfo.JobInfo)
	//	Context.SetUserProduct(r.Context(), employeeInfo.ProductMemberList)
	//}

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
	corsOptions.AllowDomain = []string{"localhost:8199", "10.24.12.84:8199", "192.168.137.1:8199", "10.80.42.65:8199", "localhost:9528", "10.80.42.65:9528", "127.0.0.1:8197"}
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
