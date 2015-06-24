package tasks

import (
	//"adexchange/engine"
	//m "adexchange/models"
	adpm "adprocess/models"
	"github.com/astaxie/beego"
	"time"
)

func DailyReportInit(minutes int) {
	timer := time.NewTicker(time.Minutes * time.Duration(minutes))
	for {
		select {
		case <-timer.C:
			beego.Debug()
			adpm.UpdateDailyReport("2015-06-18")
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
