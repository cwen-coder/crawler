package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	v := url.Values{}
	v.Add("uname", "")
	v.Add("pwd", "")
	cl := &http.Client{}

	resp, err := cl.PostForm("http://cwengo.com/admin/login", v)
	defer resp.Body.Close() //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)

}
