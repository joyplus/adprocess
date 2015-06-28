package models

import (
	"github.com/astaxie/beego/orm"
)

type PmpDemandDailyReport struct {
	Id              int    `orm:"column(id);auto"`
	AdDate          string `orm:"column(ad_date);type(date);null"`
	DemandAdspaceId int    `orm:"column(demand_adspace_id);null"`
	ReqSuccess      int    `orm:"column(req_success);null"`
	ReqTimeout      int    `orm:"column(req_timeout);null"`
	ReqNoad         int    `orm:"column(req_noad);null"`
	ReqError        int    `orm:"column(req_error);null"`
}

func init() {
	orm.RegisterModel(new(PmpDemandDailyReport))
}
