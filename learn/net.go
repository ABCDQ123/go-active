package learn

import (
	"fmt"
	"log"
	"net/http"
)

func Net_handler() {
	http.HandleFunc("/", request_parse)      // 设置访问路由地址
	err := http.ListenAndServe(":9090", nil) // 设置服务器监听端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func request_parse(w http.ResponseWriter, r *http.Request) {
	var err_r = r.ParseForm()
	if err_r != nil {
		fmt.Println("net: form erro_r")
	}
	fmt.Println("net: Form= ", r.Form)
	fmt.Println("net: path= ", r.URL.Path)
	fmt.Println("net: scheme= ", r.URL.Scheme)
	fmt.Println("net: url_long= ", r.Form["url_long"])
	//for k, v := range r.Form {
	//	fmt.Println("Key: ", k)
	//	fmt.Println("Value: ", strings.Join(v, ""))
	//}
	var _, err_w = fmt.Fprintf(w, "Hello go-active")
	if err_w != nil {
		fmt.Println("net: form erro_w")
	}
}
