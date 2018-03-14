package system

import (
	"encoding/json"
	"fmt"
	"for_job/http/lib"
	"io/ioutil"
	"log"
)

func tribOnLocal(n int) string {
	var i int
	n = n - 1
	var a, b, c uint64 = 0, 0, 1
	for i = 0; i < n; i++ {
		a, b, c = b, c, a+b+c
	}
	return fmt.Sprint(a)
}

func Config() {
	file, err := ioutil.ReadFile("config/config")
	if err != nil {
		log.Println("Error read config: ", err)
		return
	}
	err = json.Unmarshal(file, &lib.Config)
	if err != nil {
		fmt.Println("Unmarshal config", err)
	}
}
