package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

var gCurCookies []*http.Cookie
var gCurCookieJar *cookiejar.Jar

func initAll() {
	gCurCookies = nil
	//var err error;
	gCurCookieJar, _ = cookiejar.New(nil)
}

const (
	getUrl    string = "http://cwengo.com/admin/home"
	login_url string = "http://cwengo.com/admin/login"
	uname     string = "用户名"
	pwd       string = "密码"
)

func login() {
	httpClient := &http.Client{
		CheckRedirect: nil,
		Jar:           gCurCookieJar,
	}
	v := url.Values{}
	v.Add("uname", uname)
	v.Add("pwd", pwd)
	//v.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	resp, err := httpClient.PostForm(login_url, v)
	if err != nil {
		fmt.Println("login出错:", err.Error())
	}
	gCurCookies = resp.Cookies()
}
func getBody() string {
	url, _ := url.Parse(getUrl)
	gCurCookieJar.SetCookies(url, gCurCookies)
	httpClient := &http.Client{
		CheckRedirect: nil,
		Jar:           gCurCookieJar,
	}
	resp, err := httpClient.Get(getUrl)
	if err != nil {
		fmt.Println("login出错:", err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

func main() {
	initAll()
	login()
	data := getBody()
	fmt.Println(data)
}
