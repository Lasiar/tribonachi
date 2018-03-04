package lib

type DockerComposeWithoutVolumes struct {
	Version  string `yaml:"version"`
	Services struct {
		Http Web                `yaml:"http"`
		Back BackGround         `yaml:"back"`
		DB   RedisWithoutVolume `yaml:"db"`
	} `yaml:"services"`
}

type DockerComposeWithVolumes struct {
	Version  string `yaml:"version"`
	Services struct {
		Http Web             `yaml:"http"`
		Back BackGround      `yaml:"back"`
		DB   RedisWithDocker `yaml:"db"`
	} `yaml:"services"`
}

type BackGround struct {
	Build string `yaml:"build"`
}

type Web struct {
	Build string   `yaml:"build"`
	Ports []string `yaml:"ports"`
}

type RedisWithDocker struct {
	Image  string `yaml:"image"`
	Volume string `yaml:"volume"`
}

type RedisWithoutVolume struct {
	Image string `yaml:"image"`
}
