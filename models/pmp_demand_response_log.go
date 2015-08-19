package models

import (
	"github.com/astaxie/beego/orm"
)

type PmpDemandResponseLog struct {
	Id              int    `orm:"column(id);auto"`
	AdDate          string `orm:"column(ad_date);type(date);null"`
	Bid             string `orm:"column(bid);size(50);null"`
	Did             string `orm:"column(did);size(50);null"`
	ResponseTime    string `orm:"column(response_time);type(timestamp);null"`
	DemandAdspaceId int    `orm:"column(demand_adspace_id);null"`
	ResponseCode    int    `orm:"column(response_code);null"`
	ResponseBody    string `orm:"column(response_body);null"`
}

func init() {
	orm.RegisterModel(new(PmpDemandResponseLog))
}
