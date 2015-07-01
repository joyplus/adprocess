package models

import (
	"github.com/astaxie/beego/orm"
)

type PmpTrackingLog struct {
	Id              int    `orm:"column(id);auto"`
	AdDate          string `orm:"column(ad_date);type(date);null"`
	RequestTime     string `orm:"column(request_time);type(int);null"`
	Bid             string `orm:"column(bid);size(50);null"`
	LogType         int    `orm:"column(log_type);null"`
	Os              int    `orm:"column(os);null"`
	IdType          int    `orm:"column(id_type);null"`
	Pkgname         string `orm:"column(pkgname);size(50);null"`
	Uid             string `orm:"column(uid);size(30);null"`
	Ip              string `orm:"column(ip);size(20);null"`
	ProvinceCode    string `orm:"column(province_code);size(45);null"`
	CityCode        string `orm:"column(city_code);size(45);null"`
	PmpAdspaceId    int    `orm:"column(pmp_adspace_id);null"`
	DemandAdspaceId int    `orm:"column(demand_adspace_id);null"`
	Ua              string `orm:"column();size(255);null"`
}

func init() {
	orm.RegisterModel(new(PmpTrackingLog))
}
