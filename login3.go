package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

func main() {
	c1 := &http.Client{}
	req2, _ := http.NewRequest("GET", "http://cwengo.com/admin/login", nil)
	res2, _ := c1.Do(req2)

	var temp_cookies = res2.Cookies()

	for _, v := range res2.Cookies() {
		req2.AddCookie(v)
	}
	Ja, _ := cookiejar.New(nil)
	getURL1, _ := url.Parse("http://cwengo.com/admin/login")
	Ja.SetCookies(getURL1, temp_cookies)
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
	var cookies = resp.Cookies()
	c := &http.Client{}

	Jar, _ := cookiejar.New(nil)
	getURL, _ := url.Parse("http://cwengo.com/admin/home")
	Jar.SetCookies(getURL, cookies)
	c.Jar = Jar
	res, _ := c.Get("http://cwengo.com/admin/home")
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(data))

}
