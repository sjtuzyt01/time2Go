package main

import (
	"time"
	"fmt"
	"reflect"
	"strconv"
)

func main(){
	//time.time
	a := time.Now()
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))

	//string
	timeStr := a.Format("2006-01-02 15z:04:05")
	fmt.Println(timeStr)
	fmt.Println(reflect.TypeOf(timeStr))

	t, _ := time.Parse("2006-01-02 15:04:05", "2014-06-15 08:37:18")
	fmt.Println(t)
	fmt.Println(reflect.TypeOf(t))

	nDate := 20151126
	nTime := 111000000

	nYear := nDate / 10000
	fmt.Println(nYear)
	nMonth := nDate % 10000 / 100
	fmt.Println(nMonth)
	nDay := nDate % 100
	fmt.Println(nDay)
	nHour := nTime / 10000000
	fmt.Println(nHour)
	nMin := nTime / 100000 % 100
	fmt.Println(nMin)
	nSecond := nTime / 1000 % 100
	fmt.Println(nSecond)

	strYear := strconv.Itoa(nYear)
	strMonth := strconv.Itoa(nMonth)
	strDay := strconv.Itoa(nDay)
	strHour := strconv.Itoa(nHour)
	strMin := strconv.Itoa(nMin)
	strSecond := strconv.Itoa(nSecond)
	//"2014-06-15 08:37:18"
	strTime := strYear + "-" + strMonth + "-" + strDay + " " + strHour + ":" + strMin + ":" + strSecond + "0"
	fmt.Println(strTime)
	newTime, _ := time.Parse("2006-01-02 15:04:05", strTime)
	fmt.Println(newTime)
	fmt.Println(reflect.TypeOf(newTime))
}
