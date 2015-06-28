package tasks

import (
	//"adexchange/engine"
	//m "adexchange/models"
	adpm "adprocess/models"
	//"github.com/astaxie/beego"
	"time"
)

func DailyReportInit(minutes int) {
	timer := time.NewTicker(time.Minute * time.Duration(minutes))
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

func LastDayReportInit(minutes int) {
	timer := time.NewTicker(time.Minute * time.Duration(minutes))
	for {
		select {
		case <-timer.C:

		}
	}
}
