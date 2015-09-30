package models

import (
	"github.com/astaxie/beego/orm"
)

type PmpRequestLog struct {
	Id              int     `orm:"column(id);auto"`
	AdDate          string  `orm:"column(ad_date);type(date);null"`
	PmpAdspaceId    int     `orm:"column(pmp_adspace_id);null"`
	RequestTime     string  `orm:"column(request_time);type(timestamp);null"`
	Bid             string  `orm:"column(bid);size(50);null"`
	Did             string  `orm:"column(did);size(50);null"`
	Os              int     `orm:"column(os);null"`
	IdType          int     `orm:"column(id_type);null"`
	StatusCode      int     `orm:"column(status_code);null"`
	Pkgname         string  `orm:"column(pkgname);size(50);null"`
	Uid             string  `orm:"column(uid);size(30);null"`
	Ip              string  `orm:"column(ip);size(20);null"`
	ProvinceCode    string  `orm:"column(province_code);size(45);null"`
	CityCode        string  `orm:"column(city_code);size(45);null"`
	Ua              string  `orm:"column();size(255);null"`
	Lon             float32 `orm:"column(lon);null"`
	Lat             float32 `orm:"column(lat);null"`
	ProcessDuration int     `orm:"column(process_duration);null"`
	Width           int     `orm:"column(width);null"`
	Height          int     `orm:"column(height);null"`
}

func init() {
	orm.RegisterModel(new(PmpRequestLog))
}
