package web

import (
	"encoding/json"
	"fmt"
	"for_job/http/lib"
	"for_job/http/system"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var regularAlpha, _ = regexp.Compile("[[:alpha:]]")

func HttpServer(w http.ResponseWriter, r *http.Request) {
	//TODO refactor this, looks ugly
	var t lib.Response
	defer func() {
		jsBlob, _ := json.Marshal(t)
		fmt.Fprint(w, string(jsBlob))
	}()
	key := r.FormValue("n")
	if key == "" {
		t.Message = "please enter n example: '" + fmt.Sprint(r.URL)+ "?n=10'"
		return
	}
	n64, err := strconv.ParseUint(key, 10, 32)
	if err != nil {
		switch {
		case strings.Contains(fmt.Sprint(err), "parsing \"-"):
			t.Message = "error: N non positiv. N must only postiv number "
		case regularAlpha.MatchString(fmt.Sprintln(err)[28:]):
			t.Message = "error: N non number. N must only positiv number"
		default:
			t.Message = fmt.Sprint(err)
		}
		fmt.Println(fmt.Sprintln(err)[28:])
		return
	}
	n := uint32(n64)
	t.Number, err = system.Trib(n)
	if err != nil {
		t.Message = fmt.Sprint(err)
	} else {
		t.Message = "success"
	}
	return
}
