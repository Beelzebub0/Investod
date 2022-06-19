package config

import (
	"flag"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	SQL SQLConfig `yaml:"sql"`
}

type SQLConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Driver   string `yaml:"driver"`
	Database string `yaml:"database"`
	Protocol string `yaml:"protocol"`
}

func (cf *Config) GetConfig() Config {
	config := Config{}

	filename := flag.String("config", "./src/conf/conf.yaml", "Location of the config file.")
	flag.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		panic(err)
	}

	return config
}
