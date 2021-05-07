package main

import (
	"log"
	"time"
)

func main() {
	log.Println("开始时间:",time.Now().String())
	//测试代码
	a:=make(map[int]string)
	a[0]="aaaa"
	a[1]="bbbb"
	a[2]="cccc"
	a[3]="dddd"
	b:=map[int]string{
		0:"aaaa",
		1:"bbbb",
		2:"cccc",
		3:"dddd",
	}

	for i, s := range a {
		log.Println(i,s)
	}
	for i, s := range b {
		log.Println(i,s)
	}
	log.Println("结束时间:",time.Now().String())
}
