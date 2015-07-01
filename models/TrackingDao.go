package models

import (
	"adexchange/lib"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
)

func AddDemandLog(m *PmpDemandResponseLog) (err error) {

	o := orm.NewOrm()
	_, err = o.Insert(m)
	return err
}

func AddPmpTrackingLog(m *PmpTrackingLog) (err error) {

	o := orm.NewOrm()
	_, err = o.Insert(m)
	return err
}

func AddPmpRequestLog(m *PmpRequestLog) (err error) {

	o := orm.NewOrm()
	_, err = o.Insert(m)
	return err
}

func GetPmpAdspaceId(pmpAdspaceKey string) (id int) {

	id = GetCachedId("PMP_ADSPACE_" + pmpAdspaceKey)
	if id != 0 {
		return id
	}

	o := orm.NewOrm()
	pmpAdspace := PmpAdspace{PmpAdspaceKey: pmpAdspaceKey}

	err := o.Read(&pmpAdspace, "PmpAdspaceKey")

	if err == nil {
		id = pmpAdspace.Id
		SetCachedId("PMP_ADSPACE_"+pmpAdspaceKey, id)

		if err != nil {
			beego.Error(err.Error())
		}
	}
	return id
}

func GetDemandAdspaceId(adspaceKey string) (id int) {

	id = GetCachedId("DEMAND_ADSPACE_" + adspaceKey)
	if id != 0 {
		return id
	}

	o := orm.NewOrm()
	pmpDemandAdspace := PmpDemandAdspace{DemandAdspaceKey: adspaceKey}

	err := o.Read(&pmpDemandAdspace, "DemandAdspaceKey")

	if err == nil {
		id = pmpDemandAdspace.Id
		SetCachedId("DEMAND_ADSPACE_"+adspaceKey, id)

	}
	return id
}

func SetCachedId(key string, id int) {
	c := lib.Pool.Get()
	prefix := beego.AppConfig.String("runmode") + "_"

	if _, err := c.Do("SET", prefix+key, id); err != nil {
		beego.Error(err.Error())
	}

	_, err := c.Do("EXPIRE", prefix+key, 86400)
	if err != nil {
		beego.Error(err.Error())
	}

}

func GetCachedId(key string) (id int) {
	c := lib.Pool.Get()
	prefix := beego.AppConfig.String("runmode") + "_"
	id, err := redis.Int(c.Do("get", prefix+key))

	if err != nil {
		beego.Error(err.Error())
	}
	return
}
