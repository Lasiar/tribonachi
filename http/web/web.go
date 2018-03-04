package web

import (
	"encoding/json"
	"fmt"
	"for_job/http/lib"
	"for_job/http/system"
	"log"
	"net/http"
	"strconv"
)

func HttpServer(w http.ResponseWriter, r *http.Request) {
	var t lib.Tribonacci
	key := r.FormValue("n")
	n, err := strconv.Atoi(key)
	if err != nil {
		log.Println(err)
	}
	t.Number = system.Trib(n)
	jsBlob, _ := json.Marshal(t)
	fmt.Println(string(jsBlob))
	fmt.Fprint(w, string(jsBlob))
	return
}
