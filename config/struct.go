package config

var Config config

type config struct {
	DATA struct {
		Host     string `yaml:"Host" ant:"addr_ipv4"`
		Port     int    `yaml:"Port" ant:"port"`
		User     string `yaml:"User"`
		Pass     string `yaml:"Pass"`
		Database string `yaml:"Database"`
	}
	Log struct {
		Path string `yaml:"Path"`
	}
}
