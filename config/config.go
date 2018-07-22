package config

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
	"os"
)

type TypeConfig struct {
	Service string
	Hosts    map[string]string
	Ports    map[string]string
}

func LoadConfig() *TypeConfig {
	conf := &TypeConfig{}

	confPath := os.Getenv("CONFPATH")
	if confPath == "" {
		confPath = "config/api-server.yaml"
	}

	data, errReadFile := ioutil.ReadFile(confPath)
	if errReadFile != nil {
		log.Fatalf("error: %v", errReadFile)
	}

	errYaml := yaml.Unmarshal(data, &conf)
	if errYaml != nil {
		log.Fatalf("error: %v", errYaml)
	}

	prod := os.Getenv("PROD")
	if prod == "" {
		conf.Hosts["earnings-history-service"] = "127.0.0.1"
		conf.Hosts["approved-history-service"] = "127.0.0.1"
	}

	return conf
}
