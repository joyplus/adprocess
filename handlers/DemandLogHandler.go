package handlers

import (
	"adexchange/lib"
	adxm "adexchange/models"
	adpm "adprocess/models"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/vmihailenco/msgpack.v2"
	"time"
)

func HandleDemandLog() {

	c := lib.Pool.Get()
	for {
		reply, err := c.Do("brpop", beego.AppConfig.String("runmode")+"_ADMUX_DEMAND", "0")

		if err != nil {
			beego.Error(err.Error())
			continue
		}

		switch reply := reply.(type) {
		case []interface{}:
			b, _ := redis.Bytes(reply[1], nil)
			dealDemandLog(b)
			break
		case nil:

			beego.Debug("ADMUX_DEMAND Connection timeout")
			break
		default:
			beego.Debug("ADMUX_DEMAND Unknow reply:")
			beego.Debug(reply)
			break
		}

	}

	defer c.Close()
}

func dealDemandLog(b []byte) {
	var adResponse adxm.AdResponse
	err := msgpack.Unmarshal(b, &adResponse)
	if err != nil {
		beego.Error(err.Error())
	} else {
		adpm.AddDemandLog(getDemandResponseLog(&adResponse))
	}

}

func getDemandResponseLog(adResponse *adxm.AdResponse) *adpm.PmpDemandResponseLog {

	demandResponseLog := new(adpm.PmpDemandResponseLog)

	demandResponseLog.ResponseTime = time.Unix(adResponse.GetResponseTime(), 0).Format("2006-01-02 15:04:05")

	demandResponseLog.AdDate = time.Unix(adResponse.GetResponseTime(), 0).Format("2006-01-02")
	demandResponseLog.Bid = adResponse.Bid
	demandResponseLog.ResponseCode = adResponse.StatusCode
	demandResponseLog.DemandAdspaceId = adpm.GetDemandAdspaceId(adResponse.GetDemandAdspaceKey())

	return demandResponseLog
}
