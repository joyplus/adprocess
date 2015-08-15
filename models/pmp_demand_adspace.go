package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type PmpDemandAdspace struct {
	Id               int       `orm:"column(id);auto"`
	Name             string    `orm:"column(name);size(255)"`
	DelFlg           int8      `orm:"column(del_flg);null"`
	CreateUser       int       `orm:"column(create_user);null"`
	CreateTime       time.Time `orm:"column(create_time);type(timestamp);null"`
	UpdateUser       int       `orm:"column(update_user);null"`
	UpdateTime       time.Time `orm:"column(update_time);type(timestamp);null"`
	SecretKey        string    `orm:"column(secret_key);size(50);null"`
	DemandAdspaceKey string    `orm:"column(demand_adspace_key);size(50);null"`
}

func init() {
	orm.RegisterModel(new(PmpDemandAdspace))
}
