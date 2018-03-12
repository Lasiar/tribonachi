package createConfig

import (
	"encoding/json"
	"fmt"
	"for_job/configure/service"
	"io/ioutil"
	"log"
)

func New() {
	var c ConfigHttp
	fmt.Print("use default config for http? y/n")
	if service.AskForConfirmation() {
		c.SetDefaultConfig()
	} else {
		c.SetConfig()
	}
	js, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile("./http/config/config", js, 0644)
	if err != nil {
		log.Println(err)
	}
}
