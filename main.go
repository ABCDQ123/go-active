package main

import (
	"fmt"
	"go-active/learn_basic"
	"go-active/learn_http"
	"go-active/learn_server"
	"sync"
)

func init() {
	fmt.Println("*** main.go init()")
}

func main() {
	go learns_basic()
	go learns_http()
	go learns_nat()
	go learns_checker()
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}

func learns_basic() {
	learn_basic.New_const()     //变量常量
	learn_basic.New_enum()      //枚举
	learn_basic.New_array()     //数组
	learn_basic.New_slice_map() //slice
	learn_basic.New_struct()    //结构体
	learn_basic.New_interface() //接口
	learn_basic.New_goroutine() //并发
	learn_basic.New_channels()  //channel
}

func learns_http() {
	learn_http.HTTP_handler()
}

func learns_nat() {
	learn_server.Nat_Handler()
}

func learns_checker() {
	learn_server.Checker_server_loop()
}
