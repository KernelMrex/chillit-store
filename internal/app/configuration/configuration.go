package configuration

import (
	"errors"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Configuration struct {
	DB *DBConfig `yaml:"database"`
}

func NewConfig(path string) (*Configuration, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.New("[ NewConfig ] could not open file: " + err.Error())
	}
	configData, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.New("[ NewConfig ] error while reading file: " + err.Error())
	}

	config := &Configuration{}

	if err := yaml.Unmarshal(configData, config); err != nil {
		return nil, errors.New("[ NewConfig ] error while parsing config: " + err.Error())
	}

	return config, nil
}
