package util

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/lj1570693659/gfcq_tools_product/app/model"
	"github.com/lj1570693659/gfcq_tools_product/boot"
	"github.com/spf13/cast"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func GetSplitStr(norm string) []string {
	split := boot.Seg.Cut(norm, true)
	return DeleteStringSlice(split)
}

func GetStrDiff(normWords []string, realStr string) bool {
	resWords := DeleteStringSlice(boot.Seg.Cut(realStr, true))
	rate := GetStrContainsRate(normWords, resWords)
	organizeServerName := g.Config("config.toml").GetFloat64("deliverable.keyMateRate")
	return rate >= organizeServerName
}

func GetStrContainsRate(norm, realStr []string) float64 {
	findNum := 0
	for _, nv := range norm {
		for _, rv := range realStr {
			if nv == rv {
				findNum += 1
			}
		}
	}
	return FormatFloat(float64(findNum)/float64(len(norm)), 2)
}

func FormatFloat(f float64, dig int) float64 {
	result := cast.ToFloat64(strconv.FormatFloat(f, 'f', dig+1, 64))
	pow := math.Pow(10, float64(dig))
	return math.Round(result*pow) / pow
}
func CheckIsInArray(normWords []string, realStr string) bool {
	for _, v := range normWords {
		if realStr == v {
			return true
		}
	}
	return false
}

func CheckIsContainsInArray(normWords []string, realStr string) bool {
	for _, v := range normWords {
		if strings.Contains(realStr, v) {
			return true
		}
	}
	return false
}

func GetMppData(filePath, jsonFilePath string) (proArray []*model.ProductPlan, proList []*model.ProductPlan, stepNameList []string, err error) {
	pythonFile := "./packed/getMpp.py"
	cmd := exec.Command("python", pythonFile, filePath, jsonFilePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		g.Log("error").Info(output)
		return proArray, proList, stepNameList, err
	}

	// 项目任务数据
	proArray, err = getJsonFile(jsonFilePath)
	g.Log("error").Info(fmt.Sprintf("get mpp content:%v", proArray))
	proList = getTaskTree(proArray, 0)

	// 完善阀点数据
	getStepTreeName := syncTreeStepNameTask(proList, "")
	proStepArray := syncStepNameTask(proList, make([]*model.ProductPlan, 0), "")

	// 项目阀点汇总
	stepNameList = syncStepName(proList, []string{})
	return proStepArray, getStepTreeName, stepNameList, err
}

func getJsonFile(path string) ([]*model.ProductPlan, error) {
	result := make([]*model.ProductPlan, 0)
	// 读取JSON文件
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return result, err
	}

	// 解析JSON数据
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}

	return result, err
}

func getTaskTree(list []*model.ProductPlan, pId int) []*model.ProductPlan {
	var goodArr []*model.ProductPlan
	for _, v := range list {
		if v.Pid == pId {
			// 这里可以理解为每次都从最原始的数据里面找出相对就的ID进行匹配，直到找不到就返回
			child := getTaskTree(list, v.ID)
			node := &model.ProductPlan{
				ID:        v.ID,
				Name:      v.Name,
				Active:    v.Active,
				Pid:       v.Pid,
				Cid:       v.Cid,
				StartTime: gconv.Time(v.StartTime).Format("2006-01-02 15:04"),
				EndTime:   gconv.Time(v.EndTime).Format("2006-01-02 15:04"),
				Duration:  v.Duration,
				Level:     v.Level,
				Rate:      v.Rate,
				PName:     v.PName,
				Children:  child,
			}
			goodArr = append(goodArr, node)
		}
	}
	return goodArr
}

func TimeParseYYYYMMDD(in string, sub string) (out time.Time, err error) {
	layout := "2006" + sub + "01" + sub + "02"
	out, err = time.ParseInLocation(layout, in, time.Local)
	if err != nil {
		return
	}
	return
}

func syncStepNameTask(treeArray, goodArr []*model.ProductPlan, stepName string) (getStepTreeName []*model.ProductPlan) {
	for _, v := range treeArray {
		node := &model.ProductPlan{
			ID:        v.ID,
			Name:      v.Name,
			Active:    v.Active,
			Pid:       v.Pid,
			Cid:       v.Cid,
			StartTime: gconv.Time(v.StartTime).Format("2006-01-02 15:04"),
			EndTime:   gconv.Time(v.EndTime).Format("2006-01-02 15:04"),
			Duration:  v.Duration,
			Level:     v.Level,
			Rate:      v.Rate,
			PName:     v.PName,
			StepName:  v.StepName,
		}
		if v.Level < 1 {
			stepName = ""
		}
		if v.Level == 1 {
			stepName = v.Name
			v.StepName = v.Name
		}
		// 阀点层级
		if v.Level > 1 {
			v.StepName = stepName
			node.StepName = stepName
		}
		goodArr = append(goodArr, node)
		goodArr = syncStepNameTask(v.Children, goodArr, stepName)
	}
	return goodArr
}

func syncTreeStepNameTask(treeArray []*model.ProductPlan, stepName string) (getStepTreeName []*model.ProductPlan) {
	var goodArr []*model.ProductPlan
	for _, v := range treeArray {
		node := &model.ProductPlan{
			ID:        v.ID,
			Name:      v.Name,
			Active:    v.Active,
			Pid:       v.Pid,
			Cid:       v.Cid,
			StartTime: gconv.Time(v.StartTime).Format("2006-01-02 15:04"),
			EndTime:   gconv.Time(v.EndTime).Format("2006-01-02 15:04"),
			Duration:  v.Duration,
			Level:     v.Level,
			Rate:      v.Rate,
			PName:     v.PName,
			StepName:  v.StepName,
			Children:  v.Children,
		}
		if v.Level < 1 {
			stepName = ""
		}
		if v.Level == 1 {
			stepName = v.Name
			v.StepName = v.Name
		}
		// 阀点层级
		if v.Level > 1 {
			v.StepName = stepName
			node.StepName = stepName
		}
		goodArr = append(goodArr, node)
		syncTreeStepNameTask(v.Children, stepName)
	}
	return goodArr
}

func syncStepName(proList []*model.ProductPlan, stepNameList []string) (getStepName []string) {
	for _, v := range proList {
		if v.Level < 1 {
			stepNameList = syncStepName(v.Children, stepNameList)
		}
		if v.Level == 1 {
			stepNameList = append(stepNameList, v.Name)
		}
	}
	return stepNameList
}

func DeleteStringSlice(a []string) []string {
	ret := make([]string, 0)
	for _, val := range a {
		if !g.IsEmpty(val) && val != " " && val != "" {
			ret = append(ret, val)
		}
	}
	return ret
}

func DownloadFile(url, fileName string) error {
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func GetStepName(fileName string, stepList []string) string {
	if len(stepList) == 0 {
		return ""
	}
	fileNameList := make([]string, 0)
	checkStr := fileName
	if strings.Contains(fileName, " ") {
		fileNameList = strings.Split(fileName, " ")
		checkStr = fileNameList[0]
	}

	for _, v := range stepList {
		if strings.Contains(v, checkStr) {
			return v
		}
	}
	return checkStr
}
