package models

import "time"

type PmpDemandResponseLog struct {
	Id              int       `orm:"column(id);auto"`
	AdDate          string    `orm:"column(ad_date);type(date);null"`
	Bid             string    `orm:"column(bid);size(50);null"`
	ResponseTime    time.Time `orm:"column(response_time);type(timestamp);null"`
	DemandAdspaceId int       `orm:"column(demand_adspace_id);null"`
	ResponseCode    string    `orm:"column(response_code);size(10);null"`
	ResponseBody    string    `orm:"column(response_body);null"`
}
