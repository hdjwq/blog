package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)
var Pool *redis.Pool
func Init(addr string)  {
     Pool=&redis.Pool{
     	 MaxActive:100,
     	 IdleTimeout:100*time.Second,
     	 MaxIdle:10,
     	 Dial: func() (redis.Conn, error) {
			 conn,err:=redis.Dial("tcp",addr);
             return conn,err
		 },
		 TestOnBorrow: func(c redis.Conn, t time.Time) error {
			 if time.Since(t)<time.Minute {
				 return nil
			 }
			 _,err:=c.Do("PING");
			 return  err
		 },
	 }
}
func SetHash(hashName string,key interface{},val interface{})error{
	 conn:=Pool.Get()
	 _,err:=conn.Do("HSET",hashName,key,val)
	 return err
}
func GetHash(hashName string,key interface{})(result interface{},err error){
	conn:=Pool.Get()
	result,err=conn.Do("HGET",hashName,key)
	return
}
func RmHash(hashName string,key interface{})(err error){
	conn:=Pool.Get()
	_,err=conn.Do("HDEL",hashName,key)
	return
}