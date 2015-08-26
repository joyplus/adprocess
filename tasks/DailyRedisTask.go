package tasks

import (
	"adexchange/lib"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	//"time"
)

func RemoveMHQueue(strLastDay string) {

	beego.Info("Start delete mh queue for :" + beego.AppConfig.String("runmode") + "_MHQUEUE_" + strLastDay)
	c := lib.GetQueuePool.Get()

	reply, err := c.Do("keys", beego.AppConfig.String("runmode")+"_MHQUEUE_"+strLastDay+"*")

	if err != nil {
		beego.Error(err.Error())
		return
	}

	switch reply := reply.(type) {
	case []interface{}:
		for _, iKey := range reply {
			key, _ := redis.String(iKey, nil)
			beego.Info("Deleting key: " + key)
			c.Do("del", key)
		}

		break
	case nil:

		beego.Debug("ADMUX_REQ Connection timeout")
		break
	default:
		beego.Debug("ADMUX_REQ Unknow reply:")
		beego.Debug(reply)
		break
	}

	defer c.Close()
}
