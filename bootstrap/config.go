package bootstrap

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	EncKey string `yaml:"enc_key"`
	Env    string `yaml:"env"`
	Host   struct {
		Address string `yaml:"address"`
	} `yaml:"host"`
	Database struct {
		Write string `yaml:"write"`
	} `yaml:"database"`
}

func LoadConfig(file string) (cnfg Config, err error) {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return cnfg, err
	}

	err = yaml.Unmarshal([]byte(yamlFile), &cnfg)
	if err != nil {
		return cnfg, err
	}

	if cnfg.Env == "" && os.Getenv("ENV") == "" {
		cnfg.Env = "development"
	}
	if os.Getenv("ENV") != "" {
		cnfg.Env = os.Getenv("ENV")
	}
	os.Setenv("ENV", cnfg.Env)

	return cnfg, err
}
