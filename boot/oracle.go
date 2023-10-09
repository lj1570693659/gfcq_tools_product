package boot

import (
	"github.com/cengsin/oracle"
	"github.com/gogf/gf/frame/g"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var OracleDB *gorm.DB
var OracleDBName string
var CompanyObjId string
var err error

// 用于应用初始化。
func init() {
	// oracle基础信息服务
	dbLink := g.Config().GetString("database.plm.link")
	OracleDBName = g.Config().GetString("database.plm.database")
	OracleDB, err = gorm.Open(oracle.Open(dbLink), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "MBP_", // 表名前缀
			SingularTable: true,   // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{SlowThreshold: 300 * time.Millisecond}),
	})

	if err != nil {
		panic(err)
	}

	CompanyObjId = g.Config("config.toml").GetString("plm.companyObjId")
}
