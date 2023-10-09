package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/lj1570693659/gfcq_tools_product/app/service"
	"github.com/lj1570693659/gfcq_tools_product/library/response"
)

// Timer 定时任务相关
var Timer = new(timerApi)

type timerApi struct{}

// SyncProduct SignUp @summary 定时任务
// @tags    定时任务
// @produce json
// @param   entity  body model.ProductApiGetListReq true "启动定时任务"
// @router  /timer/start [GET]
// @success 200 {object} response.JsonResponse "项目清单"
func (a *timerApi) SyncProduct(r *ghttp.Request) {
	err := service.Timer.SyncProduct()
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}

}
