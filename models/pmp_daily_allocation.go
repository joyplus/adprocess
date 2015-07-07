package models

import (
	"github.com/astaxie/beego/orm"
)

type PmpDailyAllocation struct {
	Id              int     `orm:"column(id);auto"`
	AdDate          string  `orm:"column(ad_date);type(date);null"`
	PmpAdspaceId    int     `orm:"column(pmp_adspace_id);null"`
	DemandAdspaceId int     `orm:"column(demand_adspace_id);null"`
	Imp             int     `orm:"column(imp);null"`
	Clk             int     `orm:"column(clk);null"`
	Ctr             float32 `orm:"column(ctr);null"`
}

func init() {
	orm.RegisterModel(new(PmpDailyAllocation))
}
