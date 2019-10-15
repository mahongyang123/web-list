

package main

import (
    "fmt"
//    "sync"
    "time"
)

func main() {
	//fmt.Printf("start")

	toBeCharge := "16:00:00"                             //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	fmt.Println(theTime.Clock())                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	fmt.Println(time.Now().Clock())

	if time.Now().Clock().Before(theTime.Clock()) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	//fmt.Println(sr)
}