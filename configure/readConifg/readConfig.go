package readConifg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type ConfigHttp struct {
	Port   string `json:"port"`
	Router string `json:"router"`
}

func (c *ConfigHttp) ReadConfig() (err error) {
	file, err := ioutil.ReadFile("./http/config/config")
	if err != nil {
		log.Println("Error read config: ", err)
		return err
	}
	err = json.Unmarshal(file, &c)
	if err != nil {
		fmt.Println("Unmarshal config", err)
		return err
	}
	return
}
