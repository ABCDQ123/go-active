package main

import (
	"fmt"
	"go-active/learn"
)

func init() {
	fmt.Println("main.go init()")
}

func main() {
	fmt.Println("main.go main()")
	learns()
}

func learns() {
	learn.New_const()     //变量常量
	learn.New_enum()      //枚举
	learn.New_array()     //数组
	learn.New_slice_map() //slice
	learn.New_struct()    //结构体
	learn.New_interface() //接口
	learn.New_goroutine() //并发
	learn.New_channels()  //
}
