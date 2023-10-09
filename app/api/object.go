package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/lj1570693659/gfcq_tools_product/app/service"
	"github.com/lj1570693659/gfcq_tools_product/library/response"
)

// Object 员工信息API管理对象
var Object = new(objectApi)

type objectApi struct{}

// SyncObject SignUp @summary 项目清单
func (a *objectApi) SyncObject(r *ghttp.Request) {
	err := service.Object.SyncObject(r.Context())
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}
}

// SyncPlan SignUp @summary 项目清单
func (a *objectApi) SyncPlan(r *ghttp.Request) {
	err := service.Object.SyncPlan(r.Context())
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}
}

func (a *objectApi) ComputeDiff(r *ghttp.Request) {
	err := service.Object.ComputeDiff(r.Context())
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}
