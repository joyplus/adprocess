package models

import (
	"adexchange/lib"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

//adDate: 2006-01-02
func UpdateDailyReport(adDate string) (err error) {
	o := orm.NewOrm()

	beego.Info("Start update daily report")

	var records []*PmpDailyAllocation
	sql := "select allocation.id as id,allocation.ad_date as ad_date,matrix.pmp_adspace_id,matrix.demand_adspace_id from pmp_adspace_matrix as matrix inner join pmp_daily_allocation allocation on matrix.demand_adspace_id=allocation.demand_adspace_id and allocation.ad_date=? "

	paramList := []interface{}{adDate}

	_, err = o.Raw(sql, paramList).QueryRows(&records)

	if err != nil {
		return err
	}

	sql = "select count(case when log_type=2 then 1 else null end) as imp,count(case when log_type=3 then 1 else null end) as clk  from pmp_tracking_log where ad_date=? and pmp_adspace_id=? and demand_adspace_id=? "

	for _, record := range records {
		var dailyReport PmpDailyReport
		var trackingLogData PmpDailyReport
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

			_, err = o.Update(&dailyReport, "Imp", "Clk", "Ctr")
		}

		if err != nil {
			beego.Error(err.Error())
			continue
		}

		//Update allocation detail data
		if trackingLogData.Imp > 0 {
			UpdateAllocationDetail(record)
		}
	}

	return err
}

//adDate: 2006-01-02
func UpdateDemandDailyReport(adDate string) (err error) {
	o := orm.NewOrm()

	beego.Info("Start update demand daily report")

	var records []*PmpDemandDailyReport
	sql := "select distinct matrix.demand_adspace_id from pmp_adspace_matrix as matrix inner join pmp_daily_allocation allocation on matrix.demand_adspace_id=allocation.demand_adspace_id and allocation.ad_date=? "

	paramList := []interface{}{adDate}

	_, err = o.Raw(sql, paramList).QueryRows(&records)

	if err != nil {
		return err
	}

	sql = "select count(case when response_code=200 then 1 else null end) as req_success,count(case when response_code=704 then 1 else null end) as req_timeout, count(case when response_code=405 then 1 else null end) as req_noad,count(case when response_code not in(200,405,704) then 1 else null end) as req_error from pmp_demand_response_log where ad_date=? and demand_adspace_id=? "

	for _, record := range records {
		var dailyReport PmpDemandDailyReport
		var trackingLogData PmpDemandDailyReport
		paramList = []interface{}{adDate, record.DemandAdspaceId}
		err = o.Raw(sql, paramList).QueryRow(&trackingLogData)

		if err != nil {
			beego.Critical(err.Error())
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

			_, err = o.Update(&dailyReport, "ReqSuccess", "ReqTimeout", "ReqNoad", "ReqError")
		}

		if err != nil {
			beego.Error(err.Error())
			continue
		}
	}

	return err
}

//adDate: 2006-01-02
func UpdateAllocationDetail(allocation *PmpDailyAllocation) {
	o := orm.NewOrm()

	beego.Info("Start update allocation detail")

	var records []*PmpDailyAllocationDetail
	sql := "select id,allocation_id,targeting_type,targeting_code,plan_imp,plan_clk,actual_imp,actual_clk from pmp_daily_allocation_detail where allocation_id=? "

	paramList := []interface{}{allocation.Id}

	_, err := o.Raw(sql, paramList).QueryRows(&records)
	if err != nil {
		beego.Critical(err.Error())
		return
	}

	sql = "select count(case when log_type=2 then 1 else null end) as actual_imp, count(case when log_type=3 then 1 else null end) as actual_clk from pmp_tracking_log where ad_date=? and pmp_adspace_id=? and demand_adspace_id=? "
	paramList = []interface{}{allocation.AdDate, allocation.PmpAdspaceId, allocation.DemandAdspaceId, ""}
	provinceSql := sql + "and province_code=?"
	citySql := sql + "and city_code=?"

	for _, record := range records {
		var rstData *PmpDailyAllocationDetail
		if strings.EqualFold(record.TargetingType, "PROVINCE") {
			paramList[3] = record.TargetingCode
			err = o.Raw(provinceSql, paramList).QueryRow(&rstData)
		} else if strings.EqualFold(record.TargetingType, "CITY") {
			paramList[3] = record.TargetingCode
			err = o.Raw(citySql, paramList).QueryRow(&rstData)
		} else {
			continue
		}

		if err != nil {
			beego.Critical(err.Error())
			continue
		}
		if rstData != nil {
			record.ActualImp = rstData.ActualImp
			record.ActualClk = rstData.ActualClk
			_, err = o.Update(record, "ActualImp", "ActualClk")

			if err != nil {
				beego.Critical(err.Error())
				continue
			}
		}
	}

}

//adDate: 2006-01-02
func UpdateRequestDailyReport(adDate string) (err error) {
	o := orm.NewOrm()

	beego.Info("Start update request daily report")

	var records []*PmpDailyRequestReport
	sql := "select distinct matrix.pmp_adspace_id from pmp_adspace_matrix as matrix inner join pmp_daily_allocation allocation on matrix.demand_adspace_id=allocation.demand_adspace_id and allocation.ad_date=? "

	paramList := []interface{}{adDate}

	_, err = o.Raw(sql, paramList).QueryRows(&records)

	if err != nil {
		return err
	}

	sql = "select count(case when status_code=200 then 1 else null end) as req_success, count(case when status_code=405 then 1 else null end) as req_noad,count(case when status_code not in(200,405) then 1 else null end) as req_error from pmp_request_log where ad_date=? and pmp_adspace_id=? "

	for _, record := range records {
		var dailyReport PmpDailyRequestReport
		var requestLogData PmpDailyRequestReport
		paramList = []interface{}{adDate, record.PmpAdspaceId}
		err = o.Raw(sql, paramList).QueryRow(&requestLogData)

		if err != nil {
			beego.Critical(err.Error())
			continue
		}

		dailyReport = PmpDailyRequestReport{AdDate: adDate}
		dailyReport.PmpAdspaceId = record.PmpAdspaceId

		err = o.Read(&dailyReport, "AdDate", "PmpAdspaceId")

		if err == orm.ErrNoRows {
			//Tracking data
			dailyReport.ReqSuccess = requestLogData.ReqSuccess
			dailyReport.ReqNoad = requestLogData.ReqNoad
			dailyReport.ReqError = requestLogData.ReqError

			dailyReport.FillRate = lib.DivisionInt(dailyReport.ReqSuccess, dailyReport.ReqSuccess+dailyReport.ReqNoad+dailyReport.ReqError)

			_, err = o.Insert(&dailyReport)
		} else if err == nil {
			//Tracking data
			dailyReport.ReqSuccess = requestLogData.ReqSuccess
			dailyReport.ReqNoad = requestLogData.ReqNoad
			dailyReport.ReqError = requestLogData.ReqError

			dailyReport.FillRate = lib.DivisionInt(dailyReport.ReqSuccess, dailyReport.ReqSuccess+dailyReport.ReqNoad+dailyReport.ReqError)

			_, err = o.Update(&dailyReport, "ReqSuccess", "ReqNoad", "ReqError", "FillRate")
		}

		if err != nil {
			beego.Error(err.Error())
			continue
		}
	}

	return err
}

//adDate: 2006-01-02
func UpdatePmpAdspaceDailyData() (err error) {
	o := orm.NewOrm()

	beego.Info("Start update pmp adspace daily data")

	now := time.Now()
	lastDay := time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, now.Location())
	strLastDay := lastDay.Format("2006-01-02")
	last7Day := time.Date(now.Year(), now.Month(), now.Day()-8, 0, 0, 0, 0, now.Location())
	strLast7Day := last7Day.Format("2006-01-02")

	var records []*PmpDailyAllocation
	sql := "select distinct(matrix.pmp_adspace_id) from pmp_adspace_matrix as matrix inner join pmp_daily_allocation allocation on matrix.demand_adspace_id=allocation.demand_adspace_id and allocation.ad_date>=? and ?>=allocation.ad_date order by pmp_adspace_id "

	paramList := []interface{}{strLast7Day, strLastDay}

	_, err = o.Raw(sql, paramList).QueryRows(&records)

	if err != nil {
		return err
	}

	sql = "select sum(imp) as imp,sum(clk) as clk  from pmp_daily_report where ad_date>=? and ?>=ad_date and pmp_adspace_id=? "

	for _, record := range records {
		var pmpAdspace PmpAdspace
		var reportSource PmpDailyReport
		paramList = []interface{}{strLast7Day, strLastDay, record.PmpAdspaceId}
		err = o.Raw(sql, paramList).QueryRow(&reportSource)

		if err != nil {
			beego.Error(err.Error())
			continue
		}

		pmpAdspace = PmpAdspace{Id: record.PmpAdspaceId}

		err = o.Read(&pmpAdspace, "Id")

		if err == nil {
			//Tracking data
			pmpAdspace.EstDailyImp = reportSource.Imp / 7
			pmpAdspace.EstDailyClk = reportSource.Clk / 7
			pmpAdspace.EstDailyCtr = lib.DivisionInt(pmpAdspace.EstDailyClk, pmpAdspace.EstDailyImp)

			_, err = o.Update(&pmpAdspace, "EstDailyImp", "EstDailyClk", "EstDailyCtr")
		}

		if err != nil {
			beego.Error(err.Error())
			continue
		}

	}

	return err
}
