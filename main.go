package main

import (
	"adexchange/lib"
	"adprocess/handlers"
	"github.com/astaxie/beego"
)

//func main() {
//	flag.Parse()
//	pool = newPool(*redisServer, *redisPassword)

//	c := pool.Get()
//	for {
//		rstMap, _ := redis.StringMap(c.Do("brpop", "ADMUX_IMP", "0"))
//		fmt.Printf(rstMap["ADMUX_IMP"])

//	}

//	defer c.Close()

//}

//func test() {
//	b, err := msgpack.Marshal("test")
//	beego.Debug(b)
//	if err == nil {
//		c := pool.Get()
//		c.Do("lpush", "ADMUX_IMP", b)

//		defer c.Close()
//	}
//}

func main() {
	lib.Pool = lib.NewPool(beego.AppConfig.String("redis_server"), "")
	handlers.HandleImp()
}
