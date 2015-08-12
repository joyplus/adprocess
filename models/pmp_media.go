package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type PmpMedia struct {
	Id          int       `orm:"column(id);auto"`
	Name        string    `orm:"column(name);size(255)"`
	Description string    `orm:"column(description);size(500);null"`
	DelFlg      int8      `orm:"column(del_flg);null"`
	CreateUser  int       `orm:"column(create_user);null"`
	CreateTime  time.Time `orm:"column(create_time);type(timestamp);null"`
	UpdateUser  int       `orm:"column(update_user);null"`
	UpdateTime  time.Time `orm:"column(update_time);type(timestamp);null"`
}

func init() {
	orm.RegisterModel(new(PmpMedia))
}
