package boot

import (
	"github.com/robfig/cron/v3"
)

var Cron *cron.Cron

// 用于应用初始化。
func init() {
	//c := cron.New()
	//c.AddFunc("5 * * * * *", func() {
	//	fmt.Println("Every hour on the half hour-6666666666666------------", err)
	//})
	//c.Start()
	//select {}
}

//func getTimerFunc() (cron.EntryID, error) {
//	return
//}

//func Start() {
//	fmt.Println("Every hour on the half hour---------333333333333----")
//	Cron.Start()
//	fmt.Println("Every hour on the half hour---------44444444444444----")
//	for _, v := range Cron.Entries() {
//		fmt.Println("Every hour on the half hour-------------", v)
//
//	}
//	select {}
//}
