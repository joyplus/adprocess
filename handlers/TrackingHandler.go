package handlers

import (
	"adexchange/lib"
	m "adexchange/models"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/vmihailenco/msgpack.v2"
)

func HandleImp() {

	c := lib.Pool.Get()
	for {
		reply, err := c.Do("brpop", "ADMUX_IMP", "0")

		if err != nil {
			beego.Error(err.Error())
			continue
		}

		switch reply := reply.(type) {
		case []interface{}:
			b, _ := redis.Bytes(reply[1], nil)
			dealTrackingRequest(b)
		case nil:
			break
			beego.Debug("ADMUX_IMP Connection timeout")

		default:
			beego.Debug("ADMUX_IMP Unknow reply:")

			beego.Debug(reply)
			break

		}

	}

	defer c.Close()
}

func dealTrackingRequest(b []byte) {
	var adRequest m.AdRequest
	err := msgpack.Unmarshal(b, &adRequest)
	if err != nil {
		beego.Error(err.Error())
	} else {
		beego.Debug(adRequest)
	}

}
