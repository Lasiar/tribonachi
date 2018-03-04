package main

import (
	"flag"
	compose "for_job/configure/createCompose"
	conf "for_job/configure/createConfig"
	dockerFile "for_job/configure/createDockerFile"
	"runtime/debug"
)

var (
	dockerCompose *bool
	configHTTP    *bool
	docker        *bool
)

func init() {
	dockerCompose = flag.Bool("compose", false, "configure docker-compose.yml")
	configHTTP = flag.Bool("http-config", false, "configure http")
	docker = flag.Bool("docker", false, "new dockerFile")
	flag.Parse()
}

func main() {
	debug.SetGCPercent(-1)
	switch {
	case *docker:
		dockerFile.New()
	case *dockerCompose:
		compose.New()
	case *configHTTP:
		conf.New()
	default:
		conf.New()
		compose.New()
		dockerFile.New()
	}
}
