package util

import (
	"context"
	"fmt"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/crypto/gsha1"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/lj1570693659/gfcq_product_kpi/consts"
	v1 "github.com/lj1570693659/gfcq_protoc/common/v1"
	inspirit "github.com/lj1570693659/gfcq_protoc/config/inspirit/v1"
	product "github.com/lj1570693659/gfcq_protoc/config/product/v1"
	"strconv"
	"strings"
)

const (
	GSHA1 = "gsha1"
	MD5   = "md5"
)

func GetListWithPage(query *gdb.Model, page, size int32) (*gdb.Model, int32, int32, int32, error) {
	if g.IsEmpty(page) {
		page = 1
	}
	if g.IsEmpty(size) {
		size = 10
	}
	totalSize, err := query.Count()
	if err != nil {
		return query, gconv.Int32(totalSize), page, size, err
	}

	query = query.Limit(gconv.Int((page-1)*size), gconv.Int(size))
	return query, gconv.Int32(totalSize), page, size, nil
}

func Encrypt(str string) string {
	var encryptStr string
	types, _ := g.Config("config.toml").Get(context.Background(), "user.encrypt")
	switch types.String() {
	case GSHA1:
		encryptStr = gsha1.Encrypt(str)
	case MD5:
		encryptStr, _ = gmd5.Encrypt(str)
	}
	return encryptStr
}

func CheckIn(lists []uint, id uint) bool {
	isIn := false
	if len(lists) == 0 {
		return isIn
	}
	for _, v := range lists {
		if id == v {
			return true
		}
	}
	return isIn
}

func DeleteIntSlice(a []string) []string {
	ret := make([]string, 0, len(a))
	for _, val := range a {
		if !g.IsEmpty(val) {
			ret = append(ret, val)
		}
	}
	return ret
}

func DeleteInt32Slice(a []int32) []int32 {
	ret := make([]int32, 0, len(a))
	for _, val := range a {
		if !g.IsEmpty(val) {
			ret = append(ret, val)
		}
	}
	return ret
}

func GetEmploySex(name v1.SexEnum) string {
	attributeName := map[v1.SexEnum]string{
		v1.SexEnum_unknow: "未知",
		v1.SexEnum_man:    "男",
		v1.SexEnum_woman:  "女",
	}
	return attributeName[name]
}

func GetArith(name string) inspirit.ArithEnum {
	attributeName := map[string]inspirit.ArithEnum{
		"gt":  inspirit.ArithEnum_gt,
		"lt":  inspirit.ArithEnum_lt,
		"egt": inspirit.ArithEnum_egt,
		"elt": inspirit.ArithEnum_elt,
		"eq":  inspirit.ArithEnum_eq,
		"neq": inspirit.ArithEnum_neq,
	}
	return attributeName[name]
}

// GetEmployStatus  '在职状态（1：在职 2：试用期 3：实习期 4：已离职）',
func GetEmployStatus(name v1.StatusEnum) string {
	attributeName := map[v1.StatusEnum]string{
		v1.StatusEnum_anything:   "未知",
		v1.StatusEnum_working:    "在职",
		v1.StatusEnum_tryout:     "试用期",
		v1.StatusEnum_interns:    "实习期",
		v1.StatusEnum_terminated: "已离职",
	}
	return attributeName[name]
}

func GetEmployAttribute(name string) uint {
	attributeName := map[string]uint{
		"兼职": consts.PartTime,
		"全职": consts.FullTime,
		"参与": consts.PitchTime,
	}
	return attributeName[name]
}

func GetEmployAttributeId(name uint) string {
	attributeName := map[uint]string{
		consts.PartTime:  "兼职",
		consts.FullTime:  "全职",
		consts.PitchTime: "参与",
	}
	return attributeName[name]
}

func GetIsGuide(name uint) string {
	isGuideName := map[uint]string{
		0: "否",
		1: "是",
	}
	return isGuideName[name]
}

func GetFloatKeyType(name string) uint {
	devote := consts.ElseDevote
	//1：加班贡献 2：解决问题贡献 3：其他事件贡献
	if strings.Contains(name, "加班") {
		devote = consts.OverTimeDevote
	} else if strings.Contains(name, "解决问题") {
		devote = consts.SolveProblemDevote
	}
	return gconv.Uint(devote)
}

// GetFloatKeyProperty 1：正向激励 2：有待提高
func GetFloatKeyProperty(floatRaio float64) uint {
	var devote uint
	if floatRaio > 0 {
		devote = consts.ForwardDirection
	} else {
		devote = consts.ReverseDirection
	}
	return devote
}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

func DecimalLong(value float64, len int) float64 {
	formatStr := fmt.Sprintf("%%.%df", len)
	value, _ = strconv.ParseFloat(fmt.Sprintf(formatStr, value), 64)
	return value
}

// Letter 遍历a-z
func Letter(length int) []string {
	var str []string
	for i := 0; i < length; i++ {
		str = append(str, string(rune('A'+i)))
	}
	return str
}

func GetHoursIndexByScore(lists []*inspirit.CrewHoursIndexInfo, score float32) uint32 {
	for _, v := range lists {
		if get := getIndexByScore(v.ScoreRange, v.ScoreMin, v.ScoreMax, score); get {
			return v.ScoreIndex
		}
	}
	return 0
}

func GetKpiRuleByScore(lists []*inspirit.CrewKpiRuleInfo, score uint) uint {
	for _, v := range lists {
		if get := getIndexByScore(v.ScoreRange, gconv.Float32(v.ScoreMin), gconv.Float32(v.ScoreMax), gconv.Float32(score)); get {
			return gconv.Uint(v.Id)
		}
	}
	return 0
}

func GetLevelAssessByScore(lists []*inspirit.BudgetAssessInfo, score uint32) *inspirit.BudgetAssessInfo {
	for _, v := range lists {
		if get := getIndexByScore(v.ScoreRange, gconv.Float32(v.ScoreMin), gconv.Float32(v.ScoreMax), gconv.Float32(score)); get {
			return v
		}
	}
	return &inspirit.BudgetAssessInfo{}
}

func GetLevelConfirmByScore(lists []*product.LevelConfirmInfo, score uint) *product.LevelConfirmInfo {
	for _, v := range lists {
		if get := getIndexByScore(v.ScoreRange, gconv.Float32(v.ScoreMin), gconv.Float32(v.ScoreMax), gconv.Float32(score)); get {
			return v
		}
	}
	return &product.LevelConfirmInfo{}
}

func getIndexByScore(scoreRange product.ScoreRangeEnum, scoreMin, scoreMax, score float32) bool {
	switch scoreRange {
	case consts.ScoreRangeMin:
		// 左闭右开
		if scoreMin <= score && score < scoreMax {
			return true
		}
	case consts.ScoreRangeMax:
		// 左开右闭
		if scoreMin < score && score <= scoreMax {
			return true
		}
	case consts.ScoreRangeMinAndMax:
		// 左闭右闭
		if scoreMin <= score && score <= scoreMax {
			return true
		}
	case consts.NotIncludeMinMax:
		// 左右开口
		if scoreMin < score && score < scoreMax {
			return true
		}
	}
	return false
}

func GetUserRequestTypeName(methodName uint, requestModuleLists []string) string {
	requestModule := ""
	requestSecondModule := ""
	if len(requestModuleLists) > 0 {
		requestModule = requestModuleLists[0]
		requestSecondModule = requestModuleLists[1]
	}
	methodTypes := map[uint]string{
		consts.MethodGET:    "查询",
		consts.MethodPOST:   "增加",
		consts.MethodPUT:    "更新",
		consts.MethodDELETE: "删除",
	}
	moduleName := fmt.Sprintf("%s/%s", requestModule, requestSecondModule)
	secondModuleNameMap := map[string]string{
		"system/account":  "账号管理",
		"system/organize": "组织管理",
		"config/product":  "项目配置",
		"config/inspirit": "绩效配置",
		"achieve/product": "项目绩效",
		"product/create":  "项目",
		"product/delete":  "项目",
		"product/modify":  "项目",
		"product/member":  "项目成员",
		"product/stage":   "项目阀点",
	}

	// 三级模块
	third := ""
	if len(requestModuleLists) > 2 {
		switch requestModuleLists[2] {
		case "employee":
			third = "员工信息"
		case "department":
			third = "部门信息"
		case "level":
			third = "职级信息"
		case "job":
			third = "岗位信息"
		case "assess":
			third = "评级标准"
		case "confirm":
			third = "优先级"
		case "mode":
			third = "研发模式"
		case "type":
			third = "项目类型"
		case "stage":
			third = "项目阶段"
		case "roles":
			third = "项目角色"
		case "budget":
			third = "激励预算"
		case "radio":
			third = "激励应发"
		case "manage":
			third = "管理指数"
		case "hours":
			third = "工时指数"
		case "duty":
			third = "责任指数"
		case "solve":
			third = "问题解决"
		case "overtime":
			third = "加班贡献"
		case "kpiRule":
			third = "绩效等级"
		case "member":
			third = "成员绩效"
		case "crucial":
			third = "关键事件"
		case "prize":
			third = "成员奖金"
		}
	}

	actionName := ""
	switch requestModuleLists[len(requestModuleLists)-1] {
	case "export":
		actionName = "导出"
	case "import":
		actionName = "导入"
	case "compute":
		actionName = "计算"
	default:
		actionName = methodTypes[methodName]

	}
	typeName := fmt.Sprintf("%s%s%s", actionName, secondModuleNameMap[moduleName], third)
	return typeName
}
