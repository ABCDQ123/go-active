package learn_server

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

const (
	SERVER_IP      = "127.0.0.1"
	SERVER_PORT    = 9090
	SERVER_REC_LEN = 10
)

func socket_udp() {
	//create server  udp address
	address := SERVER_IP + ":" + strconv.Itoa(SERVER_PORT)
	udp_address, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println("net UDP ResolveUDPAddr err", err.Error())
		return
	}
	//listen server udp address
	udp_connect, err := net.ListenUDP("udp", udp_address)
	if err != nil {
		fmt.Println("net UDP ListenUDP err", err.Error())
		return
	}
	defer func(udp_connect *net.UDPConn) {
		err := udp_connect.Close()
		if err != nil {
			fmt.Println("net UDP Close err", err.Error())
		}
	}(udp_connect)
	//recieve data
	for true {
		data := make([]byte, SERVER_REC_LEN)
		_, udp_address_from, err := udp_connect.ReadFromUDP(data)
		if err != nil {
			fmt.Println("net UDP ReadFromUDP err", err.Error())
			return
		}
		data_string := string(data)
		data_string = strings.ToUpper(data_string)
		fmt.Println("net UDP get data :", data_string)

		data_response := "UDP did"
		err = nil
		_, err = udp_connect.WriteToUDP([]byte(data_response), udp_address_from)
		if err != nil {
			fmt.Println("net UDP WriteToUDP err", err.Error())
			return
		}
	}
}

func socket_tcp() {
	//listener server tcp address
	address := SERVER_IP + ":" + strconv.Itoa(SERVER_PORT)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("net TCP ListenTCP err", err.Error())
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("net TCP listener Close err", err.Error())
		}
	}(listener)
	//accept data
	for true {
		tcp_connect, err := listener.Accept()
		if err != nil {
			fmt.Println("net TCP Accept err", err.Error())
			return
		}
		defer func(tcp_connect net.Conn) {
			err := tcp_connect.Close()
			if err != nil {
				fmt.Println("net TCP connect Close err", err.Error())
				return
			}
		}(tcp_connect)
		for true {
			data := make([]byte, SERVER_REC_LEN)
			_, err := tcp_connect.Read(data)
			if err != nil {
				fmt.Println("net TCP Read err", err.Error())
				return
			}
			data_string := string(data)
			data_string = strings.ToUpper(data_string)
			fmt.Println("net TCP get data :", data_string)

			data_response := "TCP did"
			err = nil
			_, err = tcp_connect.Write([]byte(data_response))
			if err != nil {
				fmt.Println("net UDP WriteToUDP err", err.Error())
				return
			}
		}
	}
}
