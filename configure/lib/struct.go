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
	Build     string   `yaml:"build"`
	DependsOn []string `yaml:"depends_on"`
}

type Web struct {
	Build     string   `yaml:"build"`
	Ports     []string `yaml:"ports"`
	DependsOn []string `yaml:"depends_on"`
}

type RedisWithDocker struct {
	Image   string   `yaml:"image"`
	Volumes []string `yaml:"volumes"`
}

type RedisWithoutVolume struct {
	Image string `yaml:"image"`
}
