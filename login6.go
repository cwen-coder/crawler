package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

var gCurCookies []*http.Cookie
var gCurCookieJar *cookiejar.Jar

func initAll() {
	gCurCookies = nil
	//var err error;
	gCurCookieJar, _ = cookiejar.New(nil)
}

//get url response html
func getUrlRespHtml(url string) string {
	fmt.Printf("getUrlRespHtml, url=%s", url)

	var respHtml string = ""

	httpClient := &http.Client{
		CheckRedirect: nil,
		Jar:           gCurCookieJar,
	}

	httpReq, err := http.NewRequest("GET", url, nil)
	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		fmt.Printf("http get url=%s response error=%s\n", url, err.Error())
	}

	fmt.Printf("httpResp.Header=%s", httpResp.Header)
	fmt.Printf("httpResp.Status=%s", httpResp.Status)

	defer httpResp.Body.Close()

	body, errReadAll := ioutil.ReadAll(httpResp.Body)
	if errReadAll != nil {
		fmt.Printf("get response for url=%s got error=%s\n", url, errReadAll.Error())
	}
	//全局保存
	gCurCookies = gCurCookieJar.Cookies(httpReq.URL)

	respHtml = string(body)

	return respHtml
}

func dbgPrintCurCookies() {
	var cookieNum int = len(gCurCookies)
	fmt.Printf("cookieNum=%d", cookieNum)
	for i := 0; i < cookieNum; i++ {
		var curCk *http.Cookie = gCurCookies[i]
		fmt.Printf("\n------ Cookie [%d]------", i)
		fmt.Printf("\tName=%s", curCk.Name)
		fmt.Printf("\tValue=%s", curCk.Value)
		fmt.Printf("\tPath=%s", curCk.Path)
		fmt.Printf("\tDomain=%s", curCk.Domain)
		fmt.Printf("\tExpires=%s", curCk.Expires)
		fmt.Printf("\tRawExpires=%s", curCk.RawExpires)
		fmt.Printf("\tMaxAge=%d", curCk.MaxAge)
		fmt.Printf("\tSecure=%t", curCk.Secure)
		fmt.Printf("\tHttpOnly=%t", curCk.HttpOnly)
		fmt.Printf("\tRaw=%s", curCk.Raw)
		fmt.Printf("\tUnparsed=%s", curCk.Unparsed)
	}
}

func main() {
	initAll()

	fmt.Printf("====== step 1：get Cookie ======")
	var baiduMainUrl string = "http://www.baidu.com/"
	fmt.Printf("baiduMainUrl=%s", baiduMainUrl)
	getUrlRespHtml(baiduMainUrl)
	//dbgPrintCurCookies()

	fmt.Printf("\n====== step 2：use the Cookie ======")

	var getapiUrl string = "https://passport.baidu.com/v2/api/?getapi&class=login&tpl=mn&tangram=true"
	getUrlRespHtml(getapiUrl)
	//dbgPrintCurCookies()
}
