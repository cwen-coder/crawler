package base

import (
	"bufio"
	"bytes"
	"net/http"
	"testing"
)

func Test_NewRequest(t *testing.T) {
	re, _ := http.NewRequest("GET", "http://baidu.com", nil)
	req := NewRequest(re, 2)
	if req.HttpReq() != re {
		t.Error("HttpReq出错了")
	} else {
		t.Log("HttpReq测试通过")
	}

	if req.Depth() != 2 {
		t.Error("Depth出错了")
		t.Error(req.Depth())
	} else {
		t.Log("Depth测试通过")
	}
}

func Test_NewResponse(t *testing.T) {
	rb := bytes.NewBuffer([]byte("a string to be read"))
	r := bufio.NewReader(rb)
	rep, _ := http.ReadResponse(r, nil)
	rep1 := NewResponse(rep, 2)
	if rep1.HttpResp() != rep {
		t.Error("HttpResp出错了")
	} else {
		t.Log("HttpResp测试通过")
	}

	if rep1.Depth() != 2 {
		t.Error("Depth出错了")
		t.Error(rep1.Depth())
	} else {
		t.Log("Depth测试通过")
	}
}
