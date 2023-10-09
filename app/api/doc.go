package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/lj1570693659/gfcq_tools_product/app/model"
	"github.com/lj1570693659/gfcq_tools_product/app/service"
	"github.com/lj1570693659/gfcq_tools_product/library/response"
)

// Doc 员工信息API管理对象
var Doc = new(docApi)

type docApi struct{}

// GetList SignUp @summary 项目清单
// @tags    项目管理
// @produce json
// @param   entity  body model.ProductApiGetListReq true "项目清单"
// @router  /product/lists [GET]
// @success 200 {object} response.JsonResponse "项目清单"
func (a *docApi) GetList(r *ghttp.Request) {
	var input *model.DocListInput

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProduct, err.Error())
	}
	res, err := service.Doc.GetList(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}

}

// GetTreeTask SignUp @summary 项目筛选清单  TODO
// @tags    项目管理
// @produce json
// @param   entity  body model.ProductApiGetListReq true "项目清单"
// @router  /product/all [GET]
// @success 200 {object} response.JsonResponse "项目清单"
func (a *docApi) GetTreeTask(r *ghttp.Request) {
	var input *model.DoctListInput

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProduct, err.Error())
	}
	res, err := service.Doc.GetTreeTask(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}
}

// GetComputeByPro 对比匹配
func (a *docApi) GetComputeByPro(r *ghttp.Request) {
	var input *model.DocComputeInput

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProduct, err.Error())
	}
	res, err := service.Doc.GetComputeByPro(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}
}

// GetStatistics 查询统计数据
func (a *docApi) GetStatistics(r *ghttp.Request) {
	var input *model.DocComputeInput

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProduct, err.Error())
	}
	res, err := service.Doc.GetStatistics(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}
}

// GetProductStatistics 查询各个项目统计数据
func (a *docApi) GetProductStatistics(r *ghttp.Request) {
	res, err := service.Doc.GetProductStatistics(r.Context())
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}
}

// GetDepartStatistics 按部门统计
func (a *docApi) GetDepartStatistics(r *ghttp.Request) {
	var input *model.DocComputeInput

	if err := r.Parse(&input); err != nil {
		response.JsonExit(r, response.FormatFailProduct, err.Error())
	}

	res, err := service.Doc.GetDepartStatistics(r.Context(), input)
	if err != nil {
		response.JsonExit(r, response.GetListFailProduct, err.Error())
	} else {
		response.JsonExit(r, response.Success, "ok", res)
	}
}
