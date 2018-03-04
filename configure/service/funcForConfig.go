package service

import (
	"fmt"
	"log"
	"strconv"
)

func SetPort() string {
	var port string
	fmt.Print("port -> ")
	_, err := fmt.Scanln(&port)
	if err != nil {
		log.Fatal(err)
	}
	_, err = strconv.Atoi(port)
	if err != nil {
		log.Println("incorrect value", err, "please write the numbers")
		return SetPort()
	}
	return port
}

func SetRouter() string {
	var router string
	fmt.Print("router -> ")
	_, err := fmt.Scanln(&router)
	if err != nil {
		log.Fatal(err)
	}
	return router
}
