package uitl

import (
	"crypto/md5"
	"fmt"
	"model"
	"errors"
	"time"
	"strconv"
)
func MakeToken(id int64,) string{
	 now:=time.Now();
	 nano:=now.UnixNano()
     timeStr:=strconv.FormatInt(nano,10)
	 data:=[] byte(string(id)+timeStr)
	 has:=md5.Sum(data)
	 return fmt.Sprintf("%x",has)
}
func SysErr(result *model.Result)  {
	 result.Code=500
	 result.Data=errors.New("系统错误!")
}
func Success(result *model.Result,data interface{})  {
	result.Code=200
	result.Data=data
}