package uitl

import (
	"com.blog/modal"
	"errors"
	"crypto/md5"
	"strconv"
	"time"
	"fmt"
)
func MakeLogicErr(result *modal.ResultData,msg string)  {
	 result.Code=1002
	 result.Data=errors.New(msg)
}

func MakeToken(id int64)  string{
	 timeNow:=time.Now();
	 nano:=timeNow.UnixNano();
	 str:=strconv.FormatInt(id,10)+strconv.FormatInt(nano,10)
	 has:=md5.Sum([] byte(str))
	 return fmt.Sprintf("%x",has);
}
