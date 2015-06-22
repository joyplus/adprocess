package models

import (
	"adexchange/lib"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//adDate: 2006-01-02
func UpdateDailyReport(adDate string) (err error) {
	o := orm.NewOrm()

	beego.Debug("Start update report")

	var records []*PmpDailyReport
	sql := "select matrix.pmp_adspace_id,matrix.demand_adspace_id from pmp_adspace_matrix as matrix inner join pmp_daily_allocation allocation on matrix.demand_adspace_id=allocation.demand_adspace_id and allocation.ad_date=? "

	paramList := []interface{}{adDate}

	_, err = o.Raw(sql, paramList).QueryRows(&records)

	if err != nil {
		return err
	}

	sql = "select count(case when log_type=2 then 1 else null end) as imp,count(case when log_type=3 then 1 else null end) as clk  from pmp_tracking_log where ad_date=? and pmp_adspace_id=? and demand_adspace_id=? "
	var dailyReport PmpDailyReport
	var trackingLogData PmpDailyReport
	for _, record := range records {

		paramList = []interface{}{adDate, record.PmpAdspaceId, record.DemandAdspaceId}
		err = o.Raw(sql, paramList).QueryRow(&trackingLogData)

		if err != nil {
			beego.Error(err.Error())
			continue
		}

		dailyReport = PmpDailyReport{AdDate: adDate}
		dailyReport.PmpAdspaceId = record.PmpAdspaceId
		dailyReport.DemandAdspaceId = record.DemandAdspaceId

		err = o.Read(&dailyReport, "AdDate", "PmpAdspaceId", "DemandAdspaceId")

		if err == orm.ErrNoRows {
			//Tracking data
			dailyReport.Imp = trackingLogData.Imp
			dailyReport.Clk = trackingLogData.Clk
			dailyReport.Ctr = lib.DivisionInt(trackingLogData.Clk, trackingLogData.Imp)

			_, err = o.Insert(&dailyReport)
		} else if err == nil {
			//Tracking data
			dailyReport.Imp = trackingLogData.Imp
			dailyReport.Clk = trackingLogData.Clk
			dailyReport.Ctr = lib.DivisionInt(trackingLogData.Clk, trackingLogData.Imp)

			_, err = o.Update(&dailyReport)
		}

		if err != nil {
			beego.Error(err.Error())
			continue
		}
	}

	return err
}
