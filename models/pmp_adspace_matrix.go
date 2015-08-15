package models

import (
	"github.com/astaxie/beego/orm"
)

type PmpAdspaceMatrix struct {
	Id              int `orm:"column(id);auto"`
	PmpAdspaceId    int `orm:"column(pmp_adspace_id)"`
	DemandId        int `orm:"column(demand_id);null"`
	DemandAdspaceId int `orm:"column(demand_adspace_id);null"`
	Priority        int `orm:"column(priority);null"`
}

func init() {
	orm.RegisterModel(new(PmpAdspaceMatrix))
}
