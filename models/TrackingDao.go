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

	c := lib.Pool.Get()
	id, err := redis.Int(c.Do("get", "PMP_ADSPACE_"+pmpAdspaceKey))
	if id != 0 {
		return id
	}

	o := orm.NewOrm()
	pmpAdspace := PmpAdspace{PmpAdspaceKey: pmpAdspaceKey}

	err = o.Read(&pmpAdspace, "PmpAdspaceKey")

	if err == nil {
		id = pmpAdspace.Id
		_, err := c.Do("SET", "PMP_ADSPACE_"+pmpAdspaceKey, id)

		if err != nil {
			beego.Error(err.Error())
		}
	}
	return id
}

func GetDemandAdspaceId(adspaceKey string) (id int) {

	c := lib.Pool.Get()
	id, err := redis.Int(c.Do("get", "DEMAND_ADSPACE_"+adspaceKey))
	if id != 0 {
		return id
	}

	o := orm.NewOrm()
	pmpDemandAdspace := PmpDemandAdspace{DemandAdspaceKey: adspaceKey}

	err = o.Read(&pmpDemandAdspace, "DemandAdspaceKey")

	if err == nil {
		id = pmpDemandAdspace.Id
		_, err := c.Do("SET", "DEMAND_ADSPACE_"+adspaceKey, id)

		if err != nil {
			beego.Error(err.Error())
		}
	}
	return id
}
