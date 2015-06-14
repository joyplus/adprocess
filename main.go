package main

import (
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	flag.Parse()
	pool = newPool(*redisServer, *redisPassword)
	c := pool.Get()
	for {
		rstMap, _ := redis.StringMap(c.Do("brpop", "ADMUX_LOG", "0"))
		fmt.Printf(rstMap["ADMUX_LOG"])

	}

	defer c.Close()

}
