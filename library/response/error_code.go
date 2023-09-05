package response

const (
	/** 用户信息相关（user）错误码1开头
	* 员工信息相关（Employee）错误码2开头
	* 部门信息相关（Department）错误码3开头
	* 职级信息相关（jobLevel）错误码4开头
	* 项目信息相关（Product）错误码6开头
	* 项目组成员信息相关（Product_member）错误码61开头
	**/

	Success = 0
	// ImportFileFail 上传文件失败
	ImportFileFail     = -1
	NotSignedIn        = 101
	NotSyncEmployee    = 201
	FormatFailEmployee = 202
	CreateFailEmployee = 203
	CreateFailLog      = 204

	DepartmentFailEmployee = 3002
	JobLevelFailEmployee   = 3012
	JobFailEmployee        = 3112
	FormatFailType         = 3022
	FormatFailMode         = 3032
	FormatFailProductRoles = 3042
	CreateFailProductRoles = 3043
	ModifyFailProductRoles = 3044
	DeleteFailProductRoles = 3045
	CreateFailMode         = 3033
	ModifyFailMode         = 3034
	DeleteFailMode         = 3035
	FormatFailModeStage    = 3132
	FormatFailStageRadio   = 3052
	CreateFailStageRadio   = 3053
	ModifyFailStageRadio   = 3054
	DeleteFailStageRadio   = 3055
	FormatFailDutyIndex    = 3062
	CreateFailDutyIndex    = 3063
	ModifyFailDutyIndex    = 3064
	DeleteFailDutyIndex    = 3065
	FormatFailSolveRule    = 3072
	CreateFailSolveRule    = 3073
	ModifyFailSolveRule    = 3074
	DeleteFailSolveRule    = 3075
	FormatFailOvertimeRule = 3082
	CreateFailOvertimeRule = 3083
	ModifyFailOvertimeRule = 3084
	DeleteFailOvertimeRule = 3085
	FormatFailKpiRule      = 3092
	CreateFailKpiRule      = 3093
	ModifyFailKpiRule      = 3094
	DeleteFailKpiRule      = 3095

	// FormatFailProduct 项目基础信息格式错误
	FormatFailProduct       = 602
	CreateFailProduct       = 603
	GetListFailProduct      = 604
	GetOneFailProduct       = 605
	FormatFailProductMember = 612

	// FormatFailProductStageKpi 项目阶段绩效
	FormatFailProductStageKpi  = 6102
	CreateFailProductStageKpi  = 6103
	ModifyFailProductStageKpi  = 6104
	GetListFailProductStageKpi = 6105
	GetOneFailProductStageKpi  = 6106
	// FormatFailProductMemberKpi 项目成员绩效
	FormatFailProductMemberKpi  = 6112
	CreateFailProductMemberKpi  = 6113
	ModifyFailProductMemberKpi  = 6114
	GetListFailProductMemberKpi = 6115
	GetOneFailProductMemberKpi  = 6116
	// FormatFailProductMemberPrize 项目成员奖金
	FormatFailProductMemberPrize  = 6122
	CreateFailProductMemberPrize  = 6123
	ModifyFailProductMemberPrize  = 6124
	GetListFailProductMemberPrize = 6125
	GetOneFailProductMemberPrize  = 6126
	ExportFailProductMemberPrize  = 6127
	// FormatFailProductMemberKey 项目成员关键事件
	FormatFailProductMemberKey  = 6132
	CreateFailProductMemberKey  = 6133
	ModifyFailProductMemberKey  = 6134
	GetListFailProductMemberKey = 6135
	GetOneFailProductMemberKey  = 6136

	// FormatFailLevelAssess 配置信息
	FormatFailLevelAssess = 1002
	CreateFailLevelAssess = 1003

	FormatFailLevelConfirm = 1012
	CreateFailLevelConfirm = 1013
)
