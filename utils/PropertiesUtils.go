package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Yaml struct {
	Jwt     Jwt     `yaml:"jwt"`
	Server  Server  `yaml:"server"`
	Mongodb Mongodb `yaml:"mongodb"`
	Redis   Redis   `yaml:"redis"`
}
type Server struct {
	Port string `yaml:"port"`
}

type Jwt struct {
	Secret string `yaml:"secret"`
}

type Mongodb struct {
	Url string `yaml:"url"`
}

type Redis struct {
	Url      string `yaml:"url"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

func GetYmlProperties() *Yaml {
	c := new(Yaml)
	yamlFile, _ := ioutil.ReadFile("application.yml")
	if err := yaml.Unmarshal(yamlFile, c); err != nil {
		panic(err.Error())
	}
	return c
}
