package models

import (
	"github.com/astaxie/beego/orm"
)

type PmpDailyAllocationDetail struct {
	Id            int    `orm:"column(id);auto"`
	AllocationId  int    `orm:"column(allocation_id)"`
	TargetingType string `orm:"column(targeting_type);size(45);null"`
	TargetingCode string `orm:"column(targeting_code);size(50);null"`
	PlanImp       int    `orm:"column(plan_imp);null"`
	PlanClk       int    `orm:"column(plan_clk);null"`
	ActualImp     int    `orm:"column(actual_imp);null"`
	ActualClk     int    `orm:"column(actual_clk);null"`
}

func init() {
	orm.RegisterModel(new(PmpDailyAllocationDetail))
}
