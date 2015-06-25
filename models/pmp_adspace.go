package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type PmpAdspace struct {
	Id            int       `orm:"column(id);auto"`
	Name          string    `orm:"column(name);size(255)"`
	Description   string    `orm:"column(description);size(500);null"`
	DelFlg        int8      `orm:"column(del_flg);null"`
	CreateUser    int       `orm:"column(create_user);null"`
	CreateTime    time.Time `orm:"column(create_time);type(timestamp);null"`
	UpdateUser    int       `orm:"column(update_user);null"`
	UpdateTime    time.Time `orm:"column(update_time);type(timestamp);null"`
	PmpAdspaceKey string    `orm:"column(pmp_adspace_key);size(50);null"`
	SecretKey     string    `orm:"column(secret_key);size(50);null"`
	EstDailyImp   int       `orm:"column(est_daily_imp);null"`
	EstDailyClk   int       `orm:"column(est_daily_clk);null"`
	EstDailyCtr   float32   `orm:"column(est_daily_ctr);null"`
}

func init() {
	orm.RegisterModel(new(PmpAdspace))
}
