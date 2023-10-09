package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/yanyiwu/gojieba"
	"gopkg.in/mgo.v2"
)

// ProductMongoDB 项目库
var ProductMongoDB *mgo.Database

// UserMongoDB 用户库
var UserMongoDB *mgo.Database

// DocMongoDB 文档库
var DocMongoDB *mgo.Database
var PlmDB *mgo.Collection

var Seg *gojieba.Jieba

// 用于应用初始化。
func init() {
	// mongo基础信息服务
	mongoLink := g.Config().GetString("database.mongo.link")
	session, err := mgo.Dial(mongoLink)
	if err != nil {
		panic(err)
	}
	ProductMongoDB = session.DB("product")
	UserMongoDB = session.DB("user")
	DocMongoDB = session.DB("doc")

	// Seg
	Seg = gojieba.NewJieba()
}
