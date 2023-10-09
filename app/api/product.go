package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/lj1570693659/gfcq_tools_product/app/model"
	"github.com/lj1570693659/gfcq_tools_product/app/service"
	"github.com/lj1570693659/gfcq_tools_product/library/response"
)

// Product 员工信息API管理对象
var Product = new(productApi)

type productApi struct{}

// GetList SignUp @summary 项目清单
// @tags    项目管理
// @produce json
// @param   entity  body model.ProductApiGetListReq true "项目清单"
// @router  /product/lists [GET]
// @success 200 {object} response.JsonResponse "项目清单"
func (a *productApi) GetList(r *ghttp.Request) {
	var input *model.ProductListInput

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProduct, err.Error())
	}
	res, err := service.Product.GetList(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// GetTreeTask SignUp @summary 项目筛选清单
// @tags    项目管理
// @produce json
// @param   entity  body model.ProductApiGetListReq true "项目清单"
// @router  /product/all [GET]
// @success 200 {object} response.JsonResponse "项目清单"
func (a *productApi) GetTreeTask(r *ghttp.Request) {
	var input *model.ProductTaskListInput

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProduct, err.Error())
	}
	res, err := service.Product.GetTreeTask(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}
}

func (a *productApi) GetTreeTaskCount(r *ghttp.Request) {
	var input *model.ProductTaskListInput

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProduct, err.Error())
	}
	res, err := service.Product.GetTreeTaskCount(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}
}

func (a *productApi) SyncProduct(r *ghttp.Request) {
	err := service.Plm.SyncProduct(r.Context())
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}
}

func (a *productApi) SyncTask(r *ghttp.Request) {
	err := service.Plm.SyncTask(r.Context())
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}
}

func (a *productApi) SyncTaskDoc(r *ghttp.Request) {
	err := service.Plm.SyncTaskDoc(r.Context())
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}
}

func (a *productApi) SyncDocCount(r *ghttp.Request) {
	err := service.Plm.SyncDocCount(r.Context())
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}
}

func (a *productApi) ComputeDiff(r *ghttp.Request) {
	var input *model.ProductTaskCompute

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProduct, err.Error())
	}

	res, err := service.Plm.ComputeDiff(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}
}

func (a *productApi) Norm(r *ghttp.Request) {
	var input *model.ProductTaskCompute

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProduct, err.Error())
	}

	err := service.Deliverable.SplitNorm(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok")
	}
}
