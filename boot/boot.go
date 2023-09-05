package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
	_ "github.com/lj1570693659/gfcq_tools_product/packed"
)

// 用于应用初始化。
func init() {
	// 部门、员工基础信息服务
	//organizeServerName := g.Config("config.toml").Get("grpc.organize.link")
	//OrganizeServer, err := grpc.Dial(gconv.String(organizeServerName), grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	panic(err)
	//}
	//
	//// 公共配置服务
	//configServerName := g.Config("config.toml").Get("grpc.config.link")
	//ConfigServer, err := grpc.Dial(gconv.String(configServerName), grpc.WithTransportCredentials(insecure.NewCredentials()))

	s := g.Server("gfcq_tools_product")
	s.SetFileServerEnabled(true)
	s.AddSearchPath("./public/excel")
	s.Plugin(&swagger.Swagger{})
}
