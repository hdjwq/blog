package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)
var Pool *redis.Pool;

func Init(addr string)  {
	 Pool=newPool(addr)
}
func newPool(addr string)  *redis.Pool{
	 return &redis.Pool{
	 	   MaxActive:1000,
	 	   MaxIdle:100,
	 	   Dial: func() (redis.Conn, error) {
			    conn,err:=redis.Dial("tcp",addr);
			    return conn,err
		   },
		   TestOnBorrow: func(c redis.Conn, t time.Time) error {
			   if time.Since(t)>time.Minute {
				   _,err:=c.Do("PING")
				   return err
			   }
			   return nil
		   },
	 }
}

