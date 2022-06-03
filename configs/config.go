package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"fmt"
)

type Config interface {
	Nothing()
}

func Load(c Config, fileName string) error{
    yamlFile, err := ioutil.ReadFile(fileName)
	if err != 