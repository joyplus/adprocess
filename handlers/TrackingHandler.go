package handlers

import (
	"adexchange/lib"
	adxm "adexchange/models"
	"adexchange/tools"
	adpm "adprocess/models"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/vmihailenco/msgpack.v2"
	"time"
)

func HandleReq() {

	c := lib.Pool.Get()
	for {
		reply, err := c.Do("brpop", "ADMUX_REQ", "0")

		if err != nil {
			beego.Error(err.Error())
			continue
		}

		switch reply := reply.(type) {
		case []interface{}:
			b, _ := redis.Bytes(reply[1], nil)
			dealRequestLog(b)
			break
		case nil:

			beego.Debug("ADMUX_REQ Connection timeout")
			break
		default:
			beego.Debug("ADMUX_REQ Unknow reply:")
			beego.Debug(reply)
			break
		}

	}

	defer c.Close()
}

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
			dealTrackingLog(b, 2)
			break
		case nil:

			beego.Debug("ADMUX_IMP Connection timeout")
			break
		default:
			beego.Debug("ADMUX_IMP Unknow reply:")
			beego.Debug(reply)
			break
		}

	}

	defer c.Close()
}

func HandleClk() {

	c := lib.Pool.Get()
	for {
		reply, err := c.Do("brpop", "ADMUX_CLK", "0")

		if err != nil {
			beego.Error(err.Error())
			continue
		}

		switch reply := reply.(type) {
		case []interface{}:
			b, _ := redis.Bytes(reply[1], nil)
			dealTrackingLog(b, 3)
			break
		case nil:

			beego.Debug("ADMUX_CLK Connection timeout")
			break
		default:
			beego.Debug("ADMUX_CLK Unknow reply:")
			beego.Debug(reply)
			break
		}

	}

	defer c.Close()
}

func dealRequestLog(b []byte) {
	var adRequest adxm.AdRequest
	err := msgpack.Unmarshal(b, &adRequest)
	if err != nil {
		beego.Error(err.Error())
	} else {
		adpm.AddPmpRequestLog(getRequestLog(&adRequest))
	}

}

func dealTrackingLog(b []byte, logType int) {
	var adRequest adxm.AdRequest
	err := msgpack.Unmarshal(b, &adRequest)
	if err != nil {
		beego.Error(err.Error())
	} else {
		adpm.AddPmpTrackingLog(getTrackingLog(&adRequest, logType))
	}

}

func getRequestLog(adRequest *adxm.AdRequest) *adpm.PmpRequestLog {

	requestLog := new(adpm.PmpRequestLog)

	requestLog.RequestTime = time.Unix(adRequest.RequestTime, 0).Format("2006-01-02 15:04:05")

	requestLog.AdDate = time.Unix(adRequest.RequestTime, 0).Format("2006-01-02")
	requestLog.Bid = adRequest.Bid
	requestLog.StatusCode = adRequest.StatusCode
	requestLog.Os = adRequest.Os
	requestLog.Pkgname = adRequest.Pkgname
	idType, uid := getIdTypeAndUid(adRequest)
	requestLog.IdType = idType
	requestLog.Uid = uid
	requestLog.Ip = adRequest.Ip

	provinceCode, cityCode := tools.QueryIP(adRequest.Ip)
	requestLog.ProvinceCode = provinceCode
	requestLog.CityCode = cityCode
	requestLog.PmpAdspaceId = adpm.GetPmpAdspaceId(adRequest.AdspaceKey)

	return requestLog
}

func getTrackingLog(adRequest *adxm.AdRequest, logType int) *adpm.PmpTrackingLog {

	trackingLog := new(adpm.PmpTrackingLog)
	trackingLog.RequestTime = time.Unix(adRequest.RequestTime, 0).Format("2006-01-02 15:04:05")
	trackingLog.AdDate = time.Unix(adRequest.RequestTime, 0).Format("2006-01-02")

	trackingLog.Bid = adRequest.Bid
	trackingLog.LogType = logType
	trackingLog.Os = adRequest.Os
	trackingLog.Pkgname = adRequest.Pkgname
	idType, uid := getIdTypeAndUid(adRequest)
	trackingLog.IdType = idType
	trackingLog.Uid = uid
	trackingLog.Ip = adRequest.Ip

	provinceCode, cityCode := tools.QueryIP(adRequest.Ip)
	trackingLog.ProvinceCode = provinceCode
	trackingLog.CityCode = cityCode
	trackingLog.PmpAdspaceId = adpm.GetPmpAdspaceId(adRequest.AdspaceKey)
	trackingLog.DemandAdspaceId = adpm.GetDemandAdspaceId(adRequest.DemandAdspaceKey)

	return trackingLog
}

//0:Android￼￼￼￼￼￼￼￼
//1:iOS
//2:Windows Phone
//3:Others

//0:imei
//1:wma, 终端网卡的 MAC 地址去除冒号分隔符保持大 写
//2:aid android id
//3:aaid, android advertiser id
//4:idfa
//5:oid, ios openudid
//6:uid 非 Android、iOS 操作系统的设备唯一标识码。
func getIdTypeAndUid(adRequest *adxm.AdRequest) (idType int, uid string) {
	switch adRequest.Os {
	case 0:
		if len(adRequest.Imei) > 0 {
			idType = 0
			uid = adRequest.Imei
		} else if len(adRequest.Aid) > 0 {
			idType = 2
			uid = adRequest.Aid
		} else if len(adRequest.Aaid) > 0 {
			idType = 3
			uid = adRequest.Aaid
		}
		break
	case 1:
		if len(adRequest.Idfa) > 0 {
			idType = 4
			uid = adRequest.Idfa
		} else if len(adRequest.Oid) > 0 {
			idType = 5
			uid = adRequest.Oid
		}
		break
	case 2:
		idType = 6
		uid = adRequest.Uid
	default:
		break
	}

	return idType, uid
}
