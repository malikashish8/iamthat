package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type ProfileStruct struct {
	Profile []struct {
		Name          string
		GitUserConfig string `yaml:"git-user-config"`
		SshUserConfig string `yaml:"ssh-user-config"`
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
