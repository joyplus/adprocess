package models

import (
	"github.com/astaxie/beego/orm"
)

type PmpDailyRequestReport struct {
	Id           int     `orm:"column(id);auto"`
	AdDate       string  `orm:"column(ad_date);type(date);null"`
	PmpAdspaceId int     `orm:"column(pmp_adspace_id);null"`
	ReqSuccess   int     `orm:"column(req_success);null"`
	ReqNoad      int     `orm:"column(req_noad);null"`
	ReqError     int     `orm:"column(req_error);null"`
	FillRate     float32 `orm:"column(fill_rate);null"`
}

func init() {
	orm.RegisterModel(new(PmpDailyRequestReport))
}
