package learn_server

import (
	"fmt"
	"time"
)

func Checker_server_loop() {
	go checker_server_loop()
}

func checker_server_loop() {
	for true {
		fmt.Printf("checker_server_loop \n")
		http_living_check()
		socket_living_check()
		time.Sleep(time.Second * 5)
	}
}

func http_living_check() {

}

func socket_living_check() {

}
