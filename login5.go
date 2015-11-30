package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

var cookies []*http.Cookie

const (
	spider_base_url     string = "http://www.jikexueyuan.com/path/docker/"
	login_url           string = "http://passport.jikexueyuan.com/sso/login"
	verify_code_url     string = "http://passport.jikexueyuan.com/sso/verify"
	post_login_info_url string = "http://passport.jikexueyuan.com/submit/login?is_ajax=1"
	username            string = "用户名"
	password            string = "密码"
)

func getResultHtml(get_url string) *http.Response {
	c := &http.Client{}
	Jar, _ := cookiejar.New(nil)
	getURL, _ := url.Parse(get_url)
	Jar.SetCookies(getURL, cookies)
	c.Jar = Jar
	res, _ := c.Get(get_url)
	return res
}

func login() {
	//获取登陆界面的cookie
	c := &http.Client{}
	req, _ := http.NewRequest("GET", login_url, nil)
	res, _ := c.Do(req)

	req.URL, _ = url.Parse(verify_code_url)
	var temp_cookies = res.Cookies()

	for _, v := range res.Cookies() {
		req.AddCookie(v)
	}
	// 获取验证码
	var verify_code string
	for {
		res, _ = c.Do(req)
		file, _ := os.Create("verify.gif")
		io.Copy(file, res.Body)

		fmt.Println("请查看verify.gif， 然后输入验证码， 看不清输入0重新获取验证码")
		fmt.Scanf("%s", &verify_code)
		if verify_code != "0" {
			break
		}
		res.Body.Close()
	}
	//post数据
	postValues := url.Values{}
	postValues.Add("expire", "7")
	postValues.Add("referer", "http%3A%2F%2Fwww.jikexueyuan.com%2F")
	postValues.Add("uname", username)
	postValues.Add("password", password)
	postValues.Add("verify", verify_code)

	postURL, _ := url.Parse(post_login_info_url)
	Jar, _ := cookiejar.New(nil)
	Jar.SetCookies(postURL, temp_cookies)
	c.Jar = Jar
	res, _ = c.PostForm(post_login_info_url,
		postValues)

	cookies = res.Cookies()
	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println(string(data))
}

type DownloadBean struct {
	dirname  string
	filename string
	href     string
}

func main() {

	login()
	for _, v := range cookies {
		fmt.Println(v)
	}
	var bash_str string = "#!/bin/bash \n"
	baseMap := map[string]string{}
	doc, _ := goquery.NewDocumentFromResponse(getResultHtml(spider_base_url))
	doc.Find(".lesson-info-h2").Each(func(i int, contentSelection *goquery.Selection) {
		selection := contentSelection.Find("a")
		base_href, _ := selection.Attr("href")
		dir_name := selection.Text()
		bash_str += "mkdir \"" + dir_name + "\"\n"
		baseMap[dir_name] = base_href
		fmt.Println(dir_name, "-->", base_href)

	})
	downloadList := []DownloadBean{}

	for k, v := range baseMap {
		doc, _ := goquery.NewDocumentFromResponse(getResultHtml(v))
		doc.Find(".lessonvideo-list dd h2").Each(func(i int, contentSelection *goquery.Selection) {
			selection := contentSelection.Find("a")
			href, _ := selection.Attr("href")
			filename := selection.Text()
			fmt.Println(k, "-->", filename, "-->", href)
			bean := DownloadBean{dirname: k, href: href, filename: filename}
			downloadList = append(downloadList, bean)
		})
	}

	for _, v := range downloadList {
		doc, _ := goquery.NewDocumentFromResponse(getResultHtml(v.href))
		doc.Find("source").Each(func(i int, contentSelection *goquery.Selection) {
			download_url, _ := contentSelection.Attr("src")
			one_file := "wget " + download_url + "  -O \"./" + v.dirname + "/" + v.filename + ".mp4\"\n"
			bash_str += one_file
			fmt.Println(one_file)
		})
	}

	err := ioutil.WriteFile("./download.sh", []byte(bash_str), 0x777)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("写入脚本完成")

}
