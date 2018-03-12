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
		log.Println(err)
	}
	_, err = strconv.Atoi(port)
	if err != nil {
		log.Println("incorrect value", err, "please write the numbers")
		return SetPort()
	}
	return port
}

func SetRouter() string {
	//TODO	add validation
	var router string
	fmt.Print("router -> ")
	_, err := fmt.Scanln(&router)
	if err != nil {
		log.Fatal(err)
	}
	return router
}

func SetRedis() string {
	//TODO add validation
	var redis string
	fmt.Print("use docker-compose? (default yes) y/n -> ")
	if AskForConfirmation() {
		fmt.Print("redis url:port -> ")
		_, err := fmt.Scanln(&redis)
		if err != nil {
			log.Println(err)
		}
	} else {
		redis = "db:6379"
	}
	return redis
}
