package main

import (
	"time"
	"fmt"
	"reflect"
	"strconv"
)

func main(){
	nDate := 20151126
	nTime := 111000000
	tNow := time.Now()
	fmt.Println(tNow)
	strDate := strconv.Itoa(nDate)
	strTime := strconv.Itoa(nTime)
	strYear := strDate[0:4]
	strMonth := strDate[4:6]
	strDay := strDate[6:8]
	strHour := strTime[0:2]
	strMin := strTime[2:4]
	strSecond := strTime[4:6]
	//"2014-06-15 08:37:18"
	timeString := strYear + "-" + strMonth + "-" + strDay + " " + strHour + ":" + strMin + ":" + strSecond
	fmt.Println(timeString)
	local, _ := time.LoadLocation("Asia/Chongqing")
	newTime, _ := time.ParseInLocation("2006-01-02 15:04:05", timeString, local)
	fmt.Println(newTime)
	fmt.Println(reflect.TypeOf(newTime))
}
