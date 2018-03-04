package lib

import (
	"for_job/configure/service"
)

func (c *DockerComposeWithVolumes) SetVersion() {
	c.Version = "3"
}

func (c *DockerComposeWithoutVolumes) SetVersion() {
	c.Version = "3"
}

func (c *Web) SetHttp() {
	c.Build = "./http"
	c.Ports = append(c.Ports, service.SetPortDocer())
}

func (c *BackGround) SetBack() {
	c.Build = "./back"
}

func (c *RedisWithoutVolume) SetRedisImage() {
	c.Image = "redis:alpine"
}

func (c *RedisWithDocker) SetRedisImage() {
	c.Image = "redis:alpine"
	c.Volume = service.SetValue()
}
