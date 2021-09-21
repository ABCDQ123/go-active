package learn_http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HTTP_handler() {
	http.HandleFunc("/", http_parse)         // Rout Address
	err := http.ListenAndServe(":8080", nil) // Port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func http_parse(w http.ResponseWriter, r *http.Request) {
	var err_r = r.ParseForm()
	if err_r != nil {
		http_server_err(w, r)
	}
	fmt.Println("net: url_path= ", r.URL.Path)
	fmt.Println("net: url_scheme= ", r.URL.Scheme)
	fmt.Println("net: url_long= ", r.Form["url_long"])
	switch r.URL.Path {
	case api_get_test:
		http_server_get_test(w, r)
	case api_post_test:
		http_server_post_test(w, r)
	case api_live_checker:
		_, err := w.Write([]byte("living"))
		if err != nil {
			return
		}
	default:
		http_server_err(w, r)
	}

}

type response_http_code struct {
	code int
	err  string
}

const api_live_checker = "/live_check"
const api_get_test = "/get_test"
const api_post_test = "/post_test"

func http_server_err(w http.ResponseWriter, r *http.Request) {
	http_server_write(w, []byte("Http Wrong"))
}

func http_server_get_test(w http.ResponseWriter, r *http.Request) {
	for k, _ := range r.Form {
		if k == "id" {
			http_server_write(w, []byte("{ \"data\": get response from http_server}"))
		} else {
			http_server_write(w, []byte("Params Wrong"))
		}
	}
}

func http_server_post_test(w http.ResponseWriter, r *http.Request) {
	var json_decoder = json.NewDecoder(r.Body)
	var json_map map[string]string
	err1 := json_decoder.Decode(&json_map)
	if err1 != nil {
		return
	}
	json_map["data"] = "post response from http_server"
	var json_response, _ = json.Marshal(json_map)
	http_server_write(w, json_response)
}

func http_server_write(w http.ResponseWriter, bytes []byte) {
	_, err := w.Write(bytes)
	if err != nil {
		return
	}
}
