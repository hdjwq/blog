package logic

import (
	"com.blog/dal/redis"
	"strconv"
)
func setToken(token string,id int64) error {
	 err:=redis.HSet("token",token,id)
	 if err!=nil {
		return err
	 }
	 err=redis.HSet("id",strconv.FormatInt(id,10),token)
	 if err!=nil {
		return err
	 }
	return nil
}
func hasToken(id int64)  error{
	 result,err:=redis.HGet("id",strconv.FormatInt(id,10))
	 if err!=nil {
		return err
	 }
	if result==nil {
		return nil
	}
	 by:=result.([] byte)
	 str:=string(by);
     err=redis.HDel("token",str)
     return err
}
