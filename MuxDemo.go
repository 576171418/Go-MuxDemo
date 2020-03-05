package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	hostname string
	port int
)

func init() {
	flag.StringVar(&hostname, "hostname", "127.0.0.1", "The hostname or IP on which the REST server will listen")
	flag.IntVar(&port, "port", 8080, "The port on which the REST server will listen")
}

func main() {
	flag.Parse()
	var address = fmt.Sprintf("%s:%d", hostname, port)
	log.Println("REST service listening on", address)

	router := mux.NewRouter().StrictSlash(true)
	router.
		HandleFunc("/api/service/get", MyGetHandler).
		Methods("GET")
	router.
		HandleFunc("/api/service/{servicename}/post", MyPostHandler).
		Methods("POST")

	err := http.ListenAndServe(address, router)
	if err != nil {
		log.Fatalln("ListenAndServe err:", err)
	}
	log.Println("Server end")
}

func MyGetHandler(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	param, _ := vals["servicename"]

	var res = map[string]string{"result": "succ", "name": param[0]}
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func MyPostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	servicename, _ := vars["servicename"]

	//POST请求处理Query参数
	//vals := r.URL.Query()
	//servicetype := vals["servicetype"][0]

	//POST请求处理Body参数
	//var req map[string]interface{}
	//body, _ := ioutil.ReadAll(r.Body)
	//json.Unmarshal(body, &req)
	//servicetype := req["servicetype"].(string)

	//POST请求处理表单参数
	r.ParseForm()
	if r.Form == nil {
		fmt.Println("获取不到表单参数")
	}
	servicetype := r.Form["servicetype"]
	fmt.Println(servicetype)

	var res = map[string]string{"result": "succ", "name": servicename, "type": servicetype[0]}
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Write(response)
}