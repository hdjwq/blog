package db

import (
	"testing"
	"fmt"
	"com.blog/modal"
)
func init()  {
	Init("root:123456@tcp(localhost:3306)/blog?charset=utf8&parseTime=true")
}
func TestFindUser(t *testing.T) {
	_,err:=FindUser("woshihedongjie")
	if err!=nil {
		t.Error(err)
		return
	}
	fmt.Println("成功!")
}

func TestFindHasUser(t *testing.T) {
	 _,err:=FindHasUser("woshihedongjie")
	if err!=nil {
		t.Error(err)
		return
	}
	fmt.Println("成功!")
}
func TestFindUserInfo(t *testing.T) {
     _,err:=FindUserInfo(2)
	if err!=nil {
		t.Error(err)
		return
	}
	fmt.Println("成功!")
}
func TestRegister(t *testing.T) {
	 account:=&modal.Account{
	 	AccountNumber:"testting",
	 	Password:"2321312312",
	 }
	 _,err:=Register(account)
	if err!=nil {
		t.Error(err)
		return
	}
	fmt.Println("成功!")
}