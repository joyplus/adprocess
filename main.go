package main

import (
	"adexchange/lib"
	m "adexchange/models"
	"adexchange/tools"
	"adprocess/handlers"
	"adprocess/tasks"
	//adpm "adprocess/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"sync"
)

//func main() {
//	flag.Parse()
//	pool = newPool(*redisServer, *redisPassword)

//	c := pool.Get()
//	for {
//		rstMap, _ := redis.StringMap(c.Do("brpop", "ADMUX_IMP", "0"))
//		fmt.Printf(rstMap["ADMUX_IMP"])

//	}

//	defer c.Close()

//}

//func test() {
//	b, err := msgpack.Marshal("test")
//	beego.Debug(b)
//	if err == nil {
//		c := pool.Get()
//		c.Do("lpush", "ADMUX_IMP", b)

//		defer c.Close()
//	}
//}

func main() {

	var w sync.WaitGroup
	w.Add(1)
	beego.Debug("Start adprocess")
	beego.SetLogger("file", `{"filename":"logs/adprocess.log"}`)
	beego.SetLogFuncCall(true)
	logLevel, _ := beego.AppConfig.Int("log_level")
	beego.SetLevel(logLevel)

	orm.Debug, _ = beego.AppConfig.Bool("orm_debug")

	consoleLogFlg, _ := beego.AppConfig.Bool("log_console_flg")
	if !consoleLogFlg {
		beego.BeeLogger.DelLogger("console")
	}

	lib.SetQueuePool(lib.NewPool(beego.AppConfig.String("redis_server_queue"), ""))
	lib.SetCachePool(lib.NewPool(beego.AppConfig.String("redis_server_cache"), ""))
	tools.Init("ip.dat")
	m.Connect()

	go handlers.HandleReq()
	go handlers.HandleImp()
	go handlers.HandleClk()
	go handlers.HandleDemandLog()
	dailyReportDuration, _ := beego.AppConfig.Int("daily_report_duration")
	go tasks.DailyDemandReportInit(dailyReportDuration)
	go tasks.DailyReportInit(dailyReportDuration)
	go tasks.DailyRequestReportInit(dailyReportDuration)
	go tasks.LastDayReportInit()
	go tasks.DailyTaskInit()

	w.Wait()
}
