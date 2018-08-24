package redis

import (
	"testing"
	"fmt"
)

func init()  {
	 Init("localhost:6379");
}

func TestHSet(t *testing.T) {
     err:=HSet("token","321321312","1");
	 if err!=nil {
		t.Error(err);
		return
	 }
	 fmt.Println("成功!");
}

func TestHGet(t *testing.T) {
	  result,err:=HGet("token","321321312");
	if err!=nil {
		t.Error(err);
		return
	}
	fmt.Printf("成功:%v",result)
}