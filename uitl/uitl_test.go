package uitl

import (
	"testing"
	"fmt"
)

func TestMakeToken(t *testing.T) {
	str:=MakeToken(24);
	if str=="" {
		t.Error("生成md5错误!")
		return
	}
    fmt.Println("成功!")
}
