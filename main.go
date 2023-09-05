package main

import (
	//_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/frame/g"
	//_ "github.com/lj1570693659/gfcq_tools_product/boot"
	_ "github.com/lj1570693659/gfcq_tools_product/router"
)

// @title       `gfcq_tools_product` 重庆赣锋项目管理工具
// @version     1.0
// @description `GfcqToolsProduct`重庆赣锋项目管理工具API接口文档。
// @schemes     http
func main() {
	g.Server().Run()
}
