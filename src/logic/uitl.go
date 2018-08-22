package logic

import (
	"model"
	"github.com/go-sql-driver/mysql"
	"errors"
	"uitl"
)

func sqlErrRegister(errs error,result *model.Result)  {
	switch errs.(type) {
	case *mysql.MySQLError:
		err:=errs.(*mysql.MySQLError)
		num:=err.Number;
		if num==1062 {
			result.Code=1002;
			result.Data=errors.New("账号已存在！")
		}else {
			result.Code=500;
			result.Data=errors.New("系统错误!")
		}
	default:
		uitl.SysErr(result)
	}
}
