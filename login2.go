package main

import (
	"fmt"
	"github.com/fanan/fetion_golang"
)

func main() {
	mobileNumber := "13888888888"
	password := "88888888"
	f := fetion.NewFetion(mobileNumber, password)
	err := f.Login()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Logout()
	f.BuildUserDb()
	users := []string{"12345678901", "98765432109"}
	msg := "Hello 世界"
	f.SendSms(msg, users)
	f.SendOneself("发送成功")
}
