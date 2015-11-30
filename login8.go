package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

var gCurCookies []*http.Cookie
var gCurCookieJar *cookiejar.Jar

func initAll() {
	gCurCookies = nil
	//var err error;
	gCurCookieJar, _ = cookiejar.New(nil)
}

const (
	verify_code_url string = "http://jwc.sut.edu.cn/ACTIONVALIDATERANDOMPICTURE.APPPROCESS"
	getUrl          string = "http://jwc.sut.edu.cn/ACTIONQUERYSTUDENT.APPPROCESS"
	login_url       string = "http://jwc.sut.edu.cn/ACTIONLOGON.APPPROCESS?mode=3"
	post_login_url  string = "http://jwc.sut.edu.cn/ACTIONLOGON.APPPROCESS?mode=4"
	uname           string = "学号"
	pwd             string = "密码"
)

// func login() {
// 	c := &http.Client{}
// 	// 获取验证码
// 	var verify_code string
// 	for {
// 		res, _ := c.Get(verify_code_url)
// 		file, _ := os.Create("verify.gif")
// 		io.Copy(file, res.Body)

// 		fmt.Println("请查看verify.gif， 然后输入验证码， 看不清输入0重新获取验证码")
// 		fmt.Scanf("%s", &verify_code)
// 		if verify_code != "0" {
// 			break
// 		}
// 		url1, _ := url.Parse(login_url)
// 		gCurCookieJar.SetCookies(url1, res.Cookies())
// 		res.Body.Close()
// 	}
// 	httpClient := &http.Client{
// 		CheckRedirect: nil,
// 		Jar:           gCurCookieJar,
// 	}
// 	v := url.Values{}
// 	v.Add("WebUserNO", uname)
// 	v.Add("Password", pwd)
// 	v.Add("Agnomen", verify_code)
// 	//v.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
// 	resp, err := httpClient.PostForm(login_url, v)
// 	if err != nil {
// 		fmt.Println("login出错:", err.Error())
// 	}
// 	gCurCookies = resp.Cookies()
// }
func login() {
	httpClient := &http.Client{
		CheckRedirect: nil,
		Jar:           gCurCookieJar,
	}
	req, _ := http.NewRequest("GET", login_url, nil)
	res, _ := httpClient.Do(req)

	req.URL, _ = url.Parse(verify_code_url)
	var temp_cookies = res.Cookies()

	for _, v := range res.Cookies() {
		req.AddCookie(v)
	}

	// 获取验证码
	var verify_code string
	for {
		res, _ = httpClient.Do(req)
		file, _ := os.Create("verify.gif")
		io.Copy(file, res.Body)

		fmt.Println("请查看verify.gif， 然后输入验证码， 看不清输入0重新获取验证码")
		fmt.Scanf("%s", &verify_code)
		if verify_code != "0" {
			break
		}
		res.Body.Close()
	}
	v := url.Values{}
	v.Add("WebUserNO", uname)
	v.Add("Password", pwd)
	v.Add("Agnomen", verify_code)
	postURL, _ := url.Parse(post_login_url)
	httpClient.Jar.SetCookies(postURL, temp_cookies)
	res, _ = httpClient.PostForm(post_login_url, v)
	gCurCookies = res.Cookies()
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
