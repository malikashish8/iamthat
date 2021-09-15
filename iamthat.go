package main

import (
	"errors"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()
var configFolder = "~/.config/iamthat"
var profilePath = "profile.yaml"

func main() {
	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
	})

	configPath := ResolvePathInConfig(configFolder, profilePath)
	profileStruct, err := ReadYaml(configPath)
	if err == nil {
		profileCount := len(profileStruct.Profile)
		log.Debugf("profileCount: %v\n", profileCount)
		if profileCount == 2 {
			checkAndSwitch(profileStruct)
		} else {
			log.Fatal("More than 2 profiles! This is not supported currently.")
		}
	} else {
		log.Error("Error with processing config")
		log.Error(err)
	}
}

// Check which of the profiles is in use and swith to the other
// If none of the profiles matches then the user is prompted to select a profile
func checkAndSwitch(profileStruct ProfileStruct) {
	currentIndex, err := check(profileStruct)
	if err == nil {
		if currentIndex == 0 {
			switchTo(profileStruct, 1)
		} else {
			switchTo(profileStruct, 0)
		}
	} else {
		log.Info("None of the configured profiles is in use currently!")
		selectedProfileIndex, err := SelectStringCLI("Select profile to switch to",
			[]string{profileStruct.Profile[0].Name, profileStruct.Profile[1].Name})
		if err == nil {
			log.Debugf("Profile selected: %v", selectedProfileIndex)
			switchTo(profileStruct, selectedProfileIndex)
		}
	}
}

func check(profileStruct ProfileStruct) (int, error) {
	for i, profile := range profileStruct.Profile {
		configFilePath := ResolvePathInConfig(configFolder, profile.SshUserConfig)
		if CheckFilesSameContent(configFilePath, "~/.ssh/config") {
			return i, nil
		}
	}
	// none matched
	return -1, errors.New("no match for profiles")
}

func switchTo(profileStruct ProfileStruct, profileIndex int) {
	gitResolved := ""
	sshResolved := ""
	if profileStruct.Profile[profileIndex].GitUserConfig != "" {
		gitResolved = ResolvePathInConfig(configFolder, profileStruct.Profile[profileIndex].GitUserConfig)
	}
	CopyFile(gitResolved, "~/.gitconfig")
	if profileStruct.Profile[profileIndex].SshUserConfig != "" {
		sshResolved = ResolvePathInConfig(configFolder, profileStruct.Profile[profileIndex].SshUserConfig)
	}
	CopyFile(sshResolved, "~/.ssh/config")
	log.Info("Switched to profile " + profileStruct.Profile[profileIndex].Name)
}
