package main

import (
	"fmt"
	"time"
	)
func main()  {
	//fmt.Println(time.Now().Format(time.UnixDate))//
	//fmt.Println(time.Now().Unix())//
	times := make([]int64, 2999999999)
	var  tm int64
	var i=0
	for i=0;i<=5;i++{
		fmt.Printf("请输入时间戳：\n")
		fmt.Scanf("%d",&tm)
		times[i] = tm
		fmt.Println(time.Unix(tm,0))
	}
}
