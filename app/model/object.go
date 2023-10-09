package model

type ObjCateList struct {
	CID      string            `json:"cid" gorm:"CID" bson:"cid"`
	CNAME    string            `json:"name" gorm:"CNAME" bson:"name"`
	CTYPE    string            `json:"type" gorm:"CTYPE" bson:"type"`
	Children []*ObjProductList `json:"children" gorm:"-" bson:"children"`
}

type ObjProductList struct {
	CID                string           `json:"cid" gorm:"CID" bson:"cid"`
	CNAME              string           `json:"name" gorm:"CNAME" bson:"name"`
	CPARENTID          string           `json:"parentid" gorm:"CPARENTID" bson:"parentid"`
	CUSESTATE          string           `json:"usestate" gorm:"CUSESTATE" bson:"usestate"`
	ShouldSubmitNumber int              `json:"shouldSubmitNumber" bson:"shouldSubmitNumber"` // 应该提交份数
	RealSubmitNumber   int              `json:"realSubmitNumber" bson:"realSubmitNumber"`     // 实际提交份数
	Children           []ObjProductInfo `json:"children" gorm:"-" bson:"children"`
	PlanList           []*ProductPlan   `json:"planList" gorm:"-" bson:"plan_list"`
	Plan               []*ProductPlan   `json:"plan" gorm:"-" bson:"plan"`
	StepList           []string         `json:"stepList" gorm:"-" bson:"step_list"`
}

type ObjCate struct {
	CID      string        `json:"cid" gorm:"CID" bson:"cid"`
	CNAME    string        `json:"name" gorm:"CNAME" bson:"name"`
	CTYPE    string        `json:"type" gorm:"CTYPE" bson:"type"`
	Children []*ObjProduct `json:"children" gorm:"-" bson:"children"`
}

type ObjProduct struct {
	CID       string           `json:"cid" gorm:"CID" bson:"cid"`
	CNAME     string           `json:"name" gorm:"CNAME" bson:"name"`
	CPARENTID string           `json:"parentid" gorm:"CPARENTID" bson:"parentid"`
	CUSESTATE string           `json:"usestate" gorm:"CUSESTATE" bson:"usestate"`
	Children  []ObjProductInfo `json:"children" gorm:"-" bson:"children"`
	PlanList  []*ProductPlan   `json:"planList" gorm:"-" bson:"plan_list"`
	Plan      []*ProductPlan   `json:"plan" gorm:"-" bson:"plan"`
	StepList  []string         `json:"stepList" gorm:"-" bson:"step_list"`
}

type ObjProductInfo struct {
	CID                string `json:"cid" gorm:"CID" bson:"cid"`
	CNAME              string `json:"name" gorm:"CNAME" bson:"name"`
	CSYMBOL            string `json:"symbol" gorm:"CSYMBOL" bson:"symbol"`                                  // 代号
	CSUFFIX            string `json:"suffix" gorm:"CSUFFIX" bson:"suffix"`                                  // 后缀
	CVERSION           string `json:"version" gorm:"CVERSION" bson:"version"`                               // 版本号
	CPUBLISHTIME       string `json:"publishtime" gorm:"CPUBLISHTIME" bson:"publishtime"`                   // 发布时间
	CISLATESTPUBLISHED string `json:"islatestpublished" gorm:"CISLATESTPUBLISHED" bson:"islatestpublished"` // 是否最新发布
	FGKBM              string `json:"dutyDepart" gorm:"FGKBM" bson:"dutydepart"`                            // 责任部门
	FGGJL              string `json:"changeLog" gorm:"FGGJL" bson:"changelog"`                              // 更改记录
	FXMJD              string `json:"productStep" gorm:"FXMJD" bson:"productstep"`                          // 项目阶段
}

type ProductPlan struct {
	ID        int            `json:"id" bson:"id"`
	Name      string         `json:"name" bson:"name"`
	Active    int            `json:"active" bson:"active"`
	StartTime string         `json:"startTime" bson:"startTime"` // 开始时间
	EndTime   string         `json:"endTime" bson:"endTime"`     // 结束时间
	Duration  string         `json:"duration" bson:"duration"`   // 工期
	Level     int            `json:"level" bson:"level"`         // 等级
	Rate      string         `json:"rate" bson:"rate"`           // 完成百分比
	Pid       int            `json:"pid" bson:"pid"`             // 父级ID
	Cid       int            `json:"cid" bson:"cid"`             // 子级数量
	PName     string         `json:"pName" bson:"p_name"`        // 阀点名称
	StepName  string         `json:"stepName" bson:"step_name"`  // 阀点名称
	Children  []*ProductPlan `json:"children" bson:"children"`
}
