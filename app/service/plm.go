package service

import (
	"context"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/lj1570693659/gfcq_tools_product/app/model"
	"github.com/lj1570693659/gfcq_tools_product/boot"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"sync"
)

// Plm 获取PLM系统数据
var Plm = plmService{}

type plmService struct{}

func (s *plmService) SyncProduct(ctx context.Context) error {
	// 查询项目经理管理项目数据
	pmLists := make([]*model.UserInfo, 0)
	err := boot.UserMongoDB.C("user").Find(nil).All(&pmLists)
	if err != nil {
		return err
	}
	userCID := make([]string, 0)
	for _, v := range pmLists {
		userCID = append(userCID, v.CID)
	}

	// 查询PLM系统中项目唯一标识
	proRoles := make([]string, 0)
	boot.OracleDB.Raw(fmt.Sprintf("%s('%s')", model.GetProductGroup, strings.Join(userCID, "','"))).Scan(&proRoles)

	// 查询PLM系统中项目数据
	proLists := make([]model.MBPPROJECT, 0)
	boot.OracleDB.Raw(fmt.Sprintf("%s('%s')", model.GetProductLists, strings.Join(proRoles, "','"))).Scan(&proLists)

	// 保存至Mongo
	err = s.insertMany(proLists)
	return err
}

func (s *plmService) insertMany(proLists []model.MBPPROJECT) error {
	if len(proLists) > 0 {
		go func() {
			for _, v := range proLists {
				boot.ProductMongoDB.C("plm").Insert(v)
			}
		}()
	}
	return nil
}

func (s *plmService) SyncTask(ctx context.Context) error {
	// 查询项目唯一标识
	proIds := make([]string, 0)
	err := boot.ProductMongoDB.C("plm").Find(nil).Distinct("cid", &proIds)
	if err != nil {
		return err
	}
	if len(proIds) > 0 {
		wg := sync.WaitGroup{}
		for _, v := range proIds {
			wg.Add(1)
			taskId := v
			// 保存项目任务清单（带上下层级逻辑关系）
			go func() {
				taskInfo := make([]*model.MbpProjectTask, 0)
				sql := fmt.Sprintf(model.GetTaskInfoByVersion, taskId)
				boot.OracleDB.Raw(sql).Scan(&taskInfo)

				// 查询PLM系统中项目任务（树形）
				taskInfo = s.getTaskTreeByProject(taskId, taskInfo)

				info := model.ProductTask{}
				selector := bson.M{"productId": taskId}
				boot.ProductMongoDB.C("task").Find(selector).One(&info)
				if g.IsEmpty(info.ProductId) {
					boot.ProductMongoDB.C("task").Insert(model.ProductTask{
						ProductId: taskId,
						TaskList:  taskInfo,
					})
				} else {
					data := bson.M{"$set": bson.M{"taskList": taskInfo}}
					err = boot.ProductMongoDB.C("task").Update(selector, data)
					if err != nil {
						g.Log("error").Error(err)
					}
				}

			}()

			go func() {
				taskInfo := make([]*model.MbpProjectTask, 0)
				sql := fmt.Sprintf(model.GetTaskInfoByVersion, taskId)
				boot.OracleDB.Raw(sql).Scan(&taskInfo)

				// 查询PLM系统中项目任务(清单列表 - 方便统计)
				taskInfo = s.getTaskByProject(taskId, taskInfo)
				selector := bson.M{"productId": taskId}
				info := model.ProductTask{}
				boot.ProductMongoDB.C("taskList").Find(selector).One(&info)
				if g.IsEmpty(info.ProductId) {
					boot.ProductMongoDB.C("taskList").Insert(model.ProductTask{
						ProductId: taskId,
						TaskList:  taskInfo,
					})
				} else {
					data := bson.M{"$set": bson.M{"taskList": taskInfo}}
					err = boot.ProductMongoDB.C("taskList").Update(selector, data)
					if err != nil {
						g.Log("error").Error(err)
					}
				}

			}()

			wg.Done()
		}
		wg.Wait()
	}

	return err
}

func (s *plmService) getTaskTreeByProject(pTaskId string, taskInfo []*model.MbpProjectTask) []*model.MbpProjectTask {
	var taskArr []*model.MbpProjectTask
	if len(taskInfo) > 0 {
		for _, v := range taskInfo {
			if v.CPARENTTASKID == pTaskId {
				task := make([]*model.MbpProjectTask, 0)
				boot.OracleDB.Raw(fmt.Sprintf(model.GetTaskInfoByVersion, v.CID)).Scan(&task)

				for _, vv := range task {
					taskFile := make([]*model.MbpProjectTaskFile, 0)
					usedTaskFile := make([]*model.MbpProjectTaskFile, 0)
					boot.OracleDB.Raw(fmt.Sprintf(model.GetFileByTask, vv.CID)).Scan(&taskFile)
					if len(taskFile) > 0 {
						for _, vt := range taskFile {
							if len(vt.CNAME) > 0 {
								usedTaskFile = append(usedTaskFile, vt)
							}
						}
					}
					vv.TaskFileCount = len(usedTaskFile)
					vv.TaskFile = usedTaskFile
				}

				v.Children = task
				taskInfo = s.getTaskTreeByProject(v.CID, task)
				taskArr = append(taskArr, v)
			}
		}
	}
	return taskArr
}

// getTaskByProject 非树状结构任务清单
func (s *plmService) getTaskByProject(pTaskId string, taskInfo []*model.MbpProjectTask) []*model.MbpProjectTask {
	var taskArr []*model.MbpProjectTask
	if len(taskInfo) > 0 {
		for _, v := range taskInfo {
			if v.CPARENTTASKID == pTaskId {
				task := make([]*model.MbpProjectTask, 0)
				boot.OracleDB.Raw(fmt.Sprintf(model.GetTaskInfoByVersion, v.CID)).Scan(&task)

				for _, vv := range task {
					taskFile := make([]*model.MbpProjectTaskFile, 0)
					usedTaskFile := make([]*model.MbpProjectTaskFile, 0)
					boot.OracleDB.Raw(fmt.Sprintf(model.GetFileByTask, vv.CID)).Scan(&taskFile)
					if len(taskFile) > 0 {
						for _, vt := range taskFile {
							if len(vt.CNAME) > 0 {
								usedTaskFile = append(usedTaskFile, vt)
							}
						}
					}
					vv.TaskFileCount = len(usedTaskFile)
					vv.TaskFile = usedTaskFile
				}

				task = s.getTaskByProject(v.CID, task)
				taskArr = append(taskArr, v)
				if len(task) > 0 {
					for _, vv := range task {
						taskArr = append(taskArr, vv)
					}
				}
			}
		}
	}
	return taskArr
}

func (s *plmService) SyncTaskDoc(ctx context.Context) error {
	lists := make([]model.ProductTask, 0)
	boot.ProductMongoDB.C("taskList").Find(nil).All(&lists)
	wg := sync.WaitGroup{}
	for _, v := range lists {
		wg.Add(1)
		go func(v model.ProductTask) {
			for _, vv := range v.TaskList {
				task := make([]*model.MbpProjectTaskFile, 0)
				taskFile := make([]*model.MbpProjectTaskFile, 0)
				boot.OracleDB.Raw(fmt.Sprintf(model.GetFileByTask, vv.CID)).Scan(&task)
				if len(task) > 0 {
					for _, vt := range task {
						if len(vt.CNAME) > 0 {
							taskFile = append(taskFile, vt)
						}
					}
				}
				vv.TaskFileCount = len(taskFile)
				vv.TaskFile = taskFile
			}
			selector := bson.M{"productId": v.ProductId}
			data := bson.M{"$set": bson.M{"taskList": v.TaskList}}
			err := boot.ProductMongoDB.C("taskList").Update(selector, data)
			if err != nil {
				g.Log("error").Error(err.Error())
			}
			wg.Done()
		}(v)
	}
	wg.Wait()
	return nil
}

func (s *plmService) SyncDocCount(ctx context.Context) error {
	lists := make([]model.ProductTask, 0)
	boot.ProductMongoDB.C("task").Find(nil).All(&lists)
	for _, v := range lists {
		infos, _ := s.syncDocTaskCount(v.TaskList, 0)
		selector := bson.M{"productId": v.ProductId}
		data := bson.M{"$set": bson.M{"taskList": infos}}
		err := boot.ProductMongoDB.C("task").Update(selector, data)
		if err != nil {
			g.Log("error").Error(err.Error())
		}
	}
	return nil
}

func (s *plmService) syncDocTaskCount(taskInfo []*model.MbpProjectTask, fileCount int) ([]*model.MbpProjectTask, int) {
	var result []*model.MbpProjectTask
	if len(taskInfo) > 0 {
		for _, v := range taskInfo {
			if len(v.Children) > 0 {
				v.TaskFileCount = 0
				v.TaskFile = make([]*model.MbpProjectTaskFile, 0)
				for _, vv := range v.Children {
					if len(vv.Children) > 0 {
						vv.TaskFileCount = 0
						vv.TaskFile = make([]*model.MbpProjectTaskFile, 0)
						for _, vvv := range vv.Children {
							if len(vvv.Children) > 0 {
								vvv.TaskFileCount = 0
								vvv.TaskFile = make([]*model.MbpProjectTaskFile, 0)
								for _, vvvv := range vvv.Children {
									vvv.TaskFileCount += vvvv.TaskFileCount
									if len(vvvv.TaskFile) > 0 {
										for _, vvvvF := range vvvv.TaskFile {
											vvv.TaskFile = append(vvv.TaskFile, vvvvF)
										}
									}
								}
							}
							vv.TaskFileCount += vvv.TaskFileCount
							if len(vvv.TaskFile) > 0 {
								for _, vvvF := range vvv.TaskFile {
									vv.TaskFile = append(vv.TaskFile, vvvF)
								}
							}
						}
					}
					v.TaskFileCount += vv.TaskFileCount
					if len(vv.TaskFile) > 0 {
						for _, vvF := range vv.TaskFile {
							v.TaskFile = append(v.TaskFile, vvF)
						}
					}
				}
			}
			result = append(result, v)
		}
	}
	return taskInfo, fileCount
}

func (s *plmService) ComputeDiff(ctx context.Context, input *model.ProductTaskCompute) (resulr model.NormDiffInfo, err error) {
	norm := make([]model.NormList, 0)
	boot.ProductMongoDB.C("deliverable").Find(nil).All(&norm)

	taskInfo := model.ProductTask{}
	err = boot.ProductMongoDB.C("task").Find(bson.M{"productId": input.CID}).One(&taskInfo)

	if len(taskInfo.TaskList) > 0 {
		for _, v := range taskInfo.TaskList {
			// 已提交交付物清单
			if v.CID == input.LevelCID {
				for _, no := range norm {
					if strings.Contains(v.CNAME, no.Name) {
						// 匹配阀点
						resulr, err = Compute.ComputeAnalysis(ctx, no.Records, v.TaskFile)
					}
				}
			}
		}
	}

	return resulr, err
}
