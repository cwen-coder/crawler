package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	v := url.Values{}
	v.Add("uname", "cwenadmin")
	v.Add("pwd", "yinchengwen321")
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://cwengo.com/admin/login", body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时
	fmt.Printf("%+v\n", req)                                                         //看下发送的结构

	resp, _ := client.Do(req) //发送
	defer resp.Body.Close()   //一定要关闭resp.Body
	//data, _ := ioutil.ReadAll(resp.Cookies())
	//fmt.Println(resp.Cookies())
	req.Header.Set("Cookies", resp.Cookies())
	req, _ = http.NewRequest("GET", "http://cwengo.com/admin/home", nil)
	resp, _ = client.Do(req) //发送
	data, _ := ioutil.ReadAll(resp.Body())
	fmt.Println(data)

}
