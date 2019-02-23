package main

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func (config *Conf) GetConfigParameters(fullConfigFileName string) *Conf {

	yamlFile, err := ioutil.ReadFile(fullConfigFileName)
	if err != nil {
		log.Errorf("Cannot read config file '%s' '#%v' ", fullConfigFileName, err)
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("YAML Unmarshal: %v", err)
	}

	return config
}
