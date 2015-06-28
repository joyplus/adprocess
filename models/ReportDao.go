package models

import (
	"adexchange/lib"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//adDate: 2006-01-02
func UpdateDailyReport(adDate string) (err error) {
	o := orm.NewOrm()

	beego.Debug("Start update daily report")

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

//adDate: 2006-01-02
func UpdateDemandDailyReport(adDate string) (err error) {
	o := orm.NewOrm()

	beego.Debug("Start update demand daily report")

	var records []*PmpDemandDailyReport
	sql := "select distinct matrix.demand_adspace_id from pmp_adspace_matrix as matrix inner join pmp_daily_allocation allocation on matrix.demand_adspace_id=allocation.demand_adspace_id and allocation.ad_date=? "

	paramList := []interface{}{adDate}

	_, err = o.Raw(sql, paramList).QueryRows(&records)

	if err != nil {
		return err
	}

	sql = "select count(case when response_code=200 then 1 else null end) as req_success,count(case when response_code=704 then 1 else null end) as req_timeout, count(case when response_code=405 then 1 else null end) as req_noad,count(case when response_code not in(200,405,704) then 1 else null end) as req_error from pmp_demand_response_log where ad_date=? and demand_adspace_id=? "
	var dailyReport PmpDemandDailyReport
	var trackingLogData PmpDemandDailyReport
	for _, record := range records {

		paramList = []interface{}{adDate, record.DemandAdspaceId}
		err = o.Raw(sql, paramList).QueryRow(&trackingLogData)

		if err != nil {
			beego.Error(err.Error())
			continue
		}

		dailyReport = PmpDemandDailyReport{AdDate: adDate}
		dailyReport.DemandAdspaceId = record.DemandAdspaceId

		err = o.Read(&dailyReport, "AdDate", "DemandAdspaceId")

		if err == orm.ErrNoRows {
			//Tracking data
			dailyReport.ReqSuccess = trackingLogData.ReqSuccess
			dailyReport.ReqTimeout = trackingLogData.ReqTimeout
			dailyReport.ReqNoad = trackingLogData.ReqNoad
			dailyReport.ReqError = trackingLogData.ReqError

			_, err = o.Insert(&dailyReport)
		} else if err == nil {
			//Tracking data
			dailyReport.ReqSuccess = trackingLogData.ReqSuccess
			dailyReport.ReqTimeout = trackingLogData.ReqTimeout
			dailyReport.ReqNoad = trackingLogData.ReqNoad
			dailyReport.ReqError = trackingLogData.ReqError

			_, err = o.Update(&dailyReport)
		}

		if err != nil {
			beego.Error(err.Error())
			continue
		}
	}

	return err
}
