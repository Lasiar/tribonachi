package createConfig

import (
	"for_job/configure/service"
)

type ConfigHttp struct {
	Port   string `json:"port"`
	Router string `json:"router"`
	Redis  string `json:"redis"`
}

func (c *ConfigHttp) SetConfig() {
	c.Port = service.SetPort()
	c.Router = service.SetRouter()
	c.Router = service.SetRedis()
}

func (c *ConfigHttp) SetDefaultConfig() {
	c.Port = ":8080"
	c.Router = "/tribonachi"
	c.Redis = "db:6379"
}