package config

var Config config

type config struct {
	DATA struct {
		Host     string `yaml:"Host"`
		Port     int    `yaml:"Port"`
		User     string `yaml:"User"`
		Pass     string `yaml:"Pass"`
		Database string `yaml:"Database"`
	}
	Log struct {
		Path string `yaml:"Path"`
	}
}
