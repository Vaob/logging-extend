package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"fmt"
)

type Config interface {
	Nothing()
}

