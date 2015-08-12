package tasks

import (
	//"adexchange/engine"
	//m "adexchange/models"
	adpm "adprocess/models"
	"github.com/astaxie/beego"
	"time"
)

func DailyReportInit(minutes int) {
	timer := time.NewTicker(time.Minute * time.Duration(minutes))
	adpm.UpdateDailyReport(time.Now().Format("2006-01-02"))

	for {
		select {
		case <-timer.C:
			adpm.UpdateDailyReport(time.Now().Format("2006-01-02"))
		}
	}
}

func DailyDemandReportInit(minutes int) {
	timer := time.NewTicker(time.Minute * time.Duration(minutes))
	for {
		select {
		case <-timer.C:
			adpm.UpdateDemandDailyReport(time.Now().Format("2006-01-02"))
		}
	}
}

func DailyRequestReportInit(minutes int) {
	timer := time.NewTicker(time.Minute * time.Duration(minutes))
	for {
		select {
		case <-timer.C:
			adpm.UpdateRequestDailyReport(time.Now().Format("2006-01-02"))
		}
	}
}

func LastDayReportInit() {
	go func() {
		for {
			beego.Info("Execute last task")
			executeLastDayTask()
			now := time.Now()

			// 计算下一个1点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 1, 0, 0, next.Location())

			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}

func executeLastDayTask() {

	now := time.Now()
	lastDay := time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, now.Location())
	strLastDay := lastDay.Format("2006-01-02")

	adpm.UpdateDailyReport(strLastDay)
	adpm.UpdateDemandDailyReport(strLastDay)
	adpm.UpdateRequestDailyReport(strLastDay)
	adpm.UpdatePmpAdspaceDailyData()
}

func DailyTaskInit() {
	dbName := beego.AppConfig.String("db_name")
	go func() {
		for {
			adpm.ProcessPartition(dbName, "pmp_tracking_log", false)
			adpm.ProcessPartition(dbName, "pmp_request_log", true)
			adpm.ProcessPartition(dbName, "pmp_demand_response_log", true)

			now := time.Now()

			// 计算下一个1点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 1, 0, 0, next.Location())

			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}
