package main

import (
	"fmt"
	"gosdkaliyun/util"
)

func main() {
	rs, _ := util.SendSms("17607016210",util.GenValidateCode(6),"SMS_205393604")
	fmt.Println(rs)
}




