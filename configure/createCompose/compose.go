package createCompose

import (
	"fmt"
	"for_job/configure/lib"
	"for_job/configure/service"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func New() {
	fmt.Print("use volume (for backup) y/n ->")
	if service.AskForConfirmation() {
		createDockerCompose()
	} else {
		createDockerComposeWithoutVolume()
	}
}

func createDockerComposeWithoutVolume() {
	var c lib.DockerComposeWithoutVolumes
	c.SetVersion()
	c.Services.DB.SetRedisImage()
	c.Services.Back.SetBack()
	c.Services.Http.SetHttp()
	createFile(c)
}

func createFile(c interface{}) {
	yamlBlob, err := yaml.Marshal(c)
	if err != nil {
		log.Println(err)
	}
	ioutil.WriteFile("./docker-compose.yml", yamlBlob, 0644)
}

func createDockerCompose() {
	var c lib.DockerComposeWithVolumes
	c.Services.Http.SetHttp()
	c.Services.Back.SetBack()
	c.Services.DB.SetRedisImage()
	createFile(c)
}
