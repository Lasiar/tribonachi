package main

import (
	"fmt"
	"for_job/http/db"
	"for_job/http/lib"
	"for_job/http/system"
	"for_job/http/web"
	"net/http"
)

func init() {
	system.Config()
	db.Connect()
}

func main() {
	http.HandleFunc(lib.Config.Router, web.HttpServer)
	err := http.ListenAndServe(lib.Config.Port, nil) // set listen port

	if err != nil {
		fmt.Println(err)
	}
}
