package system

import (
	"encoding/json"
	"fmt"
	"for_job/http/lib"
	"io/ioutil"
	"log"
	"strconv"
)

func tribOnLocal(n int) string {
	n = n - 1
	a := 0
	b := 0
	c := 1
	for i := 0; i < n; i++ {
		a, b, c = b, c, a+b+c
	}
	return strconv.Itoa(a)
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
	fmt.Println(lib.Config)
}
