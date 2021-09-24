package learn_server

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"
)

func Nat_Handler() {
	nat_clear_loop()
	http.HandleFunc("/", nat_by_http)
	err := http.ListenAndServe(":8086", nil)
	if err != nil {
		log.Fatal("Nat_Handler: ", err)
	}
}

const api_nat_set = "/nat_set"
const api_nat_get = "/nat_get"

var slice_ip_online = make(map[string]bool)

func nat_by_http(w http.ResponseWriter, r *http.Request) {
	var err_r = r.ParseForm()
	if err_r != nil {
		nat_err(w, r)
	}
	switch r.URL.Path {
	case api_nat_set:
		nat_set(w, r)
	case api_nat_get:
		nat_get(w, r)
	default:
		nat_err(w, r)
	}
}

func nat_set(w http.ResponseWriter, r *http.Request) {
	var ip_address = r.RemoteAddr
	_, err := w.Write([]byte(ip_address))
	if err != nil {
		fmt.Println("nat: error= ", ip_address)
		return
	}
	fmt.Println("nat: ip_address= ", ip_address)
	slice_ip_online[ip_address] = true
}

func nat_get(w http.ResponseWriter, r *http.Request) {

}

func nat_err(w http.ResponseWriter, r *http.Request) {
	nat_write(w, []byte("nat error"))
}

func nat_write(w http.ResponseWriter, bytes []byte) {
	_, err := w.Write(bytes)
	if err != nil {
		return
	}
}

func nat_clear_loop() {
	go func() {
		for {
			time.Sleep(time.Minute * 1)
			runtime.Gosched()
			slice_ip_online = make(map[string]bool)
		}
	}()
}
