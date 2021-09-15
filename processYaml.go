package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type ProfileStruct struct {
	Profile struct {
		ConfigTemplate []struct {
			Name   string
			ToPath string `yaml:"to-path"`
		} `yaml:"config-template"`
		Config []struct {
			Name       string
			ConfigFile []struct {
				Type     string
				FromPath string `yaml:"from-path"`
			} `yaml:"config-file"`
		}
	}
}

func ReadYaml(profilePath string) (ProfileStruct, error) {
	p := ProfileStruct{}

	resolvedProfilePath := ResolvePath(profilePath)
	data, err := os.ReadFile(resolvedProfilePath)
	if err == nil {
		log.Debug("Config:\n" + string(data))
		errUnamarshal := yaml.Unmarshal([]byte(data), &p)
		if errUnamarshal == nil {
			log.Debugf("Unmarshalled Config YAML:\n%v\n", p)
		} else {
			log.Error("Failed to Unmarshal yaml")
			err = errUnamarshal
		}
	} else {
		log.Error("Error reading profile at " + resolvedProfilePath)
	}
	return p, err
}
