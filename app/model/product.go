package model

type ProductInfo struct {
	CID             string  `bson:"cid"`
	Name            string  `bson:"cname"`            // 项目名称
	ObjclassID      string  `bson:"cobjclassid"`      //
	Version         string  `bson:"cversion"`         // 版本
	Type            int32   `bson:"ctype"`            // 类型
	Durtion         float64 `bson:"cduration"`        // 工期
	PlanBeginTime   string  `bson:"cplanbegintime"`   // 计划开始时间
	PlanEndTime     string  `bson:"cplanendtime"`     // 计划结束时间
	PlanWorkHours   string  `bson:"cplanworkhours"`   // 计划工时
	RealBeginTime   string  `bson:"crealbegintime"`   // 实际开始时间
	RealEndTime     string  `bson:"crealendtime"`     // 实际结束时间
	RealWorkLoad    string  `bson:"crealworkload"`    // 实际工作量（天）
	PercentComplete string  `bson:"cpercentcomplete"` // 完成百分比
	TaskStatus      int32   `bson:"ctaskstatus"`      // 任务状态 （1：等待，2：执行，3：完成，4：终止，5：暂停）
	State           int32   `bson:"cstate"`           // 生命周期状态
	CreateTime      string  `bson:"ccreatetime"`      // 创建时间
	LastModifyTime  string  `bson:"clastmodifytime"`  // 最后修改时间
	Remark          string  `bson:"Remark"`           // 备注
}

type ProductTask struct {
	ProductId string            `json:"productId" bson:"productId"`
	TaskList  []*MbpProjectTask `json:"taskList" bson:"taskList"`
}

type ProductTaskDoc struct {
	CID           string                `json:"CID"`
	CVERSION      string                `json:"CVERSION"`
	CNAME         string                `json:"CNAME"`
	CSYMBOL       string                `json:"CSYMBOL"`
	CVERSIONGROUP string                `json:"CVERSIONGROUP"`
	CLASSNAME     string                `json:"CLASSNAME"`
	Children      []*MbpProjectTaskFile `json:"children" gorm:"-"` // 交付物清单
	TaskFileCount int                   `json:"taskFileCount"`     // 交付物统计
}

type MbpProjectTask struct {
	CID           string                `json:"CID"`
	CPARENTTASKID string                `json:"CPARENTTASKID"`
	CNAME         string                `json:"CNAME"`
	COBJCLASSID   string                `json:"COBJCLASSID"`
	CISPROJECT    int32                 `json:"CISPROJECT"`
	CTASKSTATUS   int32                 `json:"CTASKSTATUS"`
	CVERSIONGROUP string                `json:"CVERSIONGROUP"`
	CPROJECTID    string                `json:"CPROJECTID"`
	Children      []*MbpProjectTask     `json:"children" gorm:"-"` // 子级部门信息
	TaskFileCount int                   `json:"taskFileCount"`     // 交付物统计
	TaskFile      []*MbpProjectTaskFile `json:"taskFile" gorm:"-"` // 交付物清单
}

type MbpProjectTaskFile struct {
	CID        string                `json:"CID"`
	PROJECTID  string                `json:"PROJECTID"`
	CVERSION   string                `json:"CVERSION"`
	CNAME      string                `json:"CNAME"`
	CSYMBOL    string                `json:"CSYMBOL"`
	CLASSNAME  string                `json:"CLASSNAME"`
	StepName   string                `json:"stepName" gorm:"-" bson:"step_name"`     // 阀点名称
	DutyDepart string                `json:"dutyDepart" gorm:"-" bson:"duty_depart"` // 责任部门
	Children   []*MbpProjectTaskFile `json:"children" gorm:"-"`                      // 子任务信息
}

type MBPPROJECT struct {
	CID                string `json:"CID" bson:"cid"`
	CNAME              string `json:"CNAME" bson:"cname"`
	COBJCLASSID        string `json:"COBJCLASSID" bson:"cobjclassid"`
	CISPROJECT         int32  `json:"CISPROJECT" bson:"cisproject"`
	CDURATION          int32  `json:"CDURATION" bson:"cduration"`
	CISMILESTONE       int32  `json:"CISMILESTONE" bson:"cismilestone"`
	CMILESTONETIME     string `json:"CMILESTONETIME" bson:"cmilestonetime"`
	CPLANBEGINTIME     string `json:"CPLANBEGINTIME" bson:"cplanbegintime"`
	CPLANENDTIME       string `json:"CPLANENDTIME" bson:"cplanendtime"`
	CPLANWORKHOURS     string `json:"CPLANWORKHOURS" bson:"cplanworkhours"`
	CREALBEGINTIME     string `json:"CREALBEGINTIME" bson:"crealbegintime"`
	CREALENDTIME       string `json:"CREALENDTIME" bson:"crealendtime"`
	CREALWORKLOAD      string `json:"CREALWORKLOAD" bson:"crealworkload"`
	CPERCENTCOMPLETE   int32  `json:"CPERCENTCOMPLETE" bson:"cpercentcomplete"`
	CTASKSTATUS        int32  `json:"CTASKSTATUS" bson:"ctaskstatus"`
	CISAUTOFINISH      int32  `json:"CISAUTOFINISH" bson:"cisautofinish"`
	CCREATETIME        string `json:"CCREATETIME" bson:"ccreatetime"`
	CCREATOR           string `json:"CCREATOR" bson:"ccreator"`
	CTYPE              int32  `json:"CTYPE" bson:"ctype"`
	CREMARK            string `json:"CREMARK" bson:"cremark"`
	CMODIFYER          string `json:"CMODIFYER" bson:"cmodifyer"`
	CLASTMODIFYTIME    string `json:"CLASTMODIFYTIME" bson:"clastmodifytime"`
	CMODIFYSTATE       int32  `json:"CMODIFYSTATE" bson:"cmodifystate"`
	CMODIFYUSER        string `json:"CMODIFYUSER" bson:"cmodifyuser"`
	CMILESTONEBEGIN    int32  `json:"CMILESTONEBEGIN" bson:"cmilestonebegin"`
	CMILESTONEEND      int32  `json:"CMILESTONEEND" bson:"cmilestoneend"`
	CMODIFYTASKPLAN    int32  `json:"CMODIFYTASKPLAN" bson:"cmodifytaskplan"`
	CREGIONID          string `json:"CREGIONID" bson:"cregionid"`
	CISOVERREGION      int32  `json:"CISOVERREGION" bson:"cisoverregion"`
	CVERSION           string `json:"CVERSION" bson:"cversion"`
	CINNERVERSION      int32  `json:"CINNERVERSION" bson:"cinnerversion"`
	CVERSIONGROUP      string `json:"CVERSIONGROUP" bson:"cversiongroup"`
	CPUBLISHTIME       string `json:"CPUBLISHTIME" bson:"cpublishtime"`
	CSTATE             int32  `json:"CSTATE" bson:"cstate"`
	CISLATEST          int32  `json:"CISLATEST" bson:"cislatest"`
	CISLATESTPUBLISHED int32  `json:"CISLATESTPUBLISHED" bson:"cislatestpublished"`
	CFREEZETIME        string `json:"CFREEZETIME" bson:"cfreezetime"`
	CSAFELEVEL         int32  `json:"CSAFELEVEL" bson:"csafelevel"`
}

type MBPPROJECTPARENTCHILDREL struct {
	CID                  string  `json:"CID"`
	CVERSIONGROUP        string  `json:"CVERSIONGROUP"`
	CPARENTTASKID        string  `json:"CPARENTTASKID"`
	CPARENTCLASSID       string  `json:"CPARENTCLASSID"`
	CPARENTVERSIONGROUP  string  `json:"CPARENTVERSIONGROUP"`
	CINDEX               float64 `json:"CINDEX"`
	CPROJECTID           string  `json:"CPROJECTID"`
	CPROJECTCLASSID      string  `json:"CPROJECTCLASSID"`
	CPROJECTVERSIONGROUP string  `json:"CPROJECTVERSIONGROUP"`
	CISBROWSE            int32   `json:"CISBROWSE"`
	CISMANAGER           int32   `json:"CISMANAGER"`
	CREFPROJECTID        string  `json:"CREFPROJECTID"`
}

type ProductListInput struct {
	CID   string `json:"cid" bson:"cid"`
	CNAME string `json:"cname" bson:"cname"`
	Page  int    `json:"page"`
	Size  int    `json:"size"`
}

type ProductTaskListInput struct {
	CID   string `v:"required#项目唯一标识不能为空" json:"cid" bson:"cid"`
	CNAME string `json:"cname" bson:"cname"`
}

// ProductTaskCompute 交付物比对
type ProductTaskCompute struct {
	CID      string `v:"required#项目唯一标识不能为空" json:"cid"`
	LevelCID string `v:"required#项目阀点唯一标识不能为空" json:"LevelCID"`
	CNAME    string `json:"cname" bson:"cname"`
}
