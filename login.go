package mian

import (
	"net/http"
	"net/url"
)

func login2(header map[string]interface{}) string {

	currentTime := time.Now().Unix() + int64(rand.Float32()*999)

	v := url.Values{}
	v.Set("entry", "weibo")
	v.Add("gateway", "1")
	v.Add("from", "")
	v.Add("savestate", "7")
	v.Add("useticket", "1")
	v.Add("service", "miniblog")
	v.Add("servertime", cov(header["servertime"]))
	v.Add("rsakv", cov(header["rsakv"]))
	v.Add("pcid", cov(header["pcid"]))
	v.Add("nonce", cov(header["nonce"]))
	v.Add("pwencode", "rsa2")
	v.Add("returntype", "META")
	v.Add("encoding", "UTF-8")
	v.Add("vsnf", "1")
	v.Add("pagerefer", "")
	v.Add("url", "http://weibo.com/ajaxlogin.php?framelogin=1&callback=parent.sinaSSOController.feedBackUrlCallBack")

	v.Add("su", base64.StdEncoding.EncodeToString([]byte("XXX@gmail.com")))

	password := sinaRSA2SSOEncoder(cov(header["pubkey"]), "wangshangyou", cov(header["servertime"]), cov(header["nonce"]))
	v.Add("sp", password)

	currentTime1 := time.Now().Unix() + int64(rand.Float32()*999)

	prelt := fmt.Sprintf("%f", math.Max(float64(currentTime1-currentTime), 100.0))

	prelt = strings.Trim(prelt, ".000000")

	v.Add("prelt", prelt)

	client := &http.Client{nil, nil, jar}

	reqest, err := http.NewRequest("POST", "http://login.sina.com.cn/sso/login.php?client=ssologin.js(v1.4.5)", strings.NewReader(v.Encode()))

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}

	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Add("Accept-Encoding", "gzip, deflate")
	reqest.Header.Add("Accept-Language", "zh-cn,zh;q=0.8,en-us;q=0.5,en;q=0.3")
	reqest.Header.Add("Connection", "keep-alive")
	reqest.Header.Add("Host", "login.sina.com.cn")
	reqest.Header.Add("Referer", "http://weibo.com/")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
	response, err := client.Do(reqest)
	defer response.Body.Close()

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}

	if response.StatusCode == 200 {

		var body []byte

		switch response.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(response.Body)
			body = dumpGZIP(reader)
		default:
			bodyByte, _ := ioutil.ReadAll(response.Body)
			body = bodyByte
		}

		r := regexp.MustCompile(`location.replace\("(.*?)"\)`)
		rs := r.FindStringSubmatch(string(body))

		// if strings.Contains(rs[1], "retcode=0") {

		// }

		return rs[1]
	}

	return ""
}

//reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
