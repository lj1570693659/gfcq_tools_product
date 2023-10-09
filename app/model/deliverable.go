package model

type NormList struct {
	Name    string          `bson:"name" json:"name"`
	Records []*NormFileInfo `bson:"records" json:"records"`
}

type NormAndRealDiffInfo struct {
	IsSubmit     bool                `bson:"is_submit" json:"isSubmit"`
	ShouldSubmit *NormFileInfo       `bson:"norm_file_info" json:"normFileInfo"`
	RealSubmit   *MbpProjectTaskFile `bson:"mbp_project_task_file" json:"mbpProjectTaskFile"`
}

type NormFileInfo struct {
	Index          int      `json:"index" bson:"index"`
	StepName       string   `json:"stepName" bson:"step_name"` // 阀点名称
	Activity       string   `json:"activity" bson:"activity"`
	DutyDepart     string   `json:"dutyDepart" bson:"duty_depart"`
	DutyUser       string   `json:"dutyUser" bson:"duty_user"`
	Input          string   `json:"input" bson:"input"`
	Name           string   `json:"name" bson:"name"`
	SplitName      []string `json:"splitName" bson:"splitName"`
	NextRoles      string   `json:"nextRoles" bson:"next_roles"`
	IsMust         string   `json:"isMust" bson:"is_must"`
	Norm           string   `json:"norm" bson:"norm"`
	Identify       string   `json:"identify" bson:"identify"`
	IdentifyNumber string   `json:"identifyNumber" bson:"identify_number"`
	Remark         string   `json:"remark" bson:"remark"`
}

// NormDiffInfo 交付物统计数据结构
type NormDiffInfo struct {
	CID                 string                `json:"cid" bson:"cid"`                               // 阀点唯一标识
	CName               string                `json:"name" bson:"name"`                             // 阀点名称
	ProjectID           string                `json:"projectid" bson:"projectid"`                   // 项目ID
	ChangeTime          string                `json:"changeTime" bson:"changeTime"`                 // 数据更新时间
	ShouldSubmitNumber  int                   `json:"shouldSubmitNumber" bson:"shouldSubmitNumber"` // 应该提交份数
	RealSubmitNumber    int                   `json:"realSubmitNumber" bson:"realSubmitNumber"`     // 实际提交份数
	ShouldSubmitList    []*NormFileInfo       `json:"shouldSubmitList" bson:"shouldSubmitList"`     // 应该提交交付物清单
	RealSubmitList      []*MbpProjectTaskFile `json:"realSubmitList" bson:"realSubmitList"`         // 实际提交交付物清单
	LackSubmitList      []*NormFileInfo       `json:"lackSubmitList" bson:"lackSubmitList"`         // 缺少的交付物清单
	NormAndRealDiffList []NormAndRealDiffInfo `json:"normAndRealDiffInfo" bson:"normAndRealDiffInfo"`
}

// ProductStageLint 导入项目组成员数据结构
type ProductStageLint struct {
	StageName   []string  `json:"stageName"`
	StageQuota  []float64 `json:"stageQuota"`
	StageBudget []float64 `json:"stageBudget"`
}

// ProductStatistics 项目总计
type ProductStatistics struct {
	ProductName        []string `json:"productName"`
	ShouldSubmitNumber []int    `json:"shouldSubmitNumber"`
	RealSubmitNumber   []int    `json:"realSubmitNumber"`
	LackSubmitNumber   []int    `json:"lackSubmitNumber"`
}

// DutyPartStatistics 责任部门统计
type DutyPartStatistics struct {
	StepName    string          `json:"stepName" bson:"step_name"`
	CountDetail []DutyPartCount `json:"countDetail" bson:"count_detail"`
}

type DutyPartCount struct {
	Count      int    `json:"value" bson:"count"`
	DutyDepart string `json:"name" bson:"duty_depart"`
}
