package main

import (
	"errors"
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()
var configFolder = "~/.config/iamthat"
var profilePath = "profile.yaml"
var currentProfileFile = ".current"

func main() {
	log.SetLevel(logrus.InfoLevel)
	envLogLevel, _ := os.LookupEnv("LOG_LEVEL")
	if envLogLevel == "DEBUG" {
		log.SetLevel(logrus.DebugLevel)
	}
	log.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
	})

	configPath := ResolvePathInConfig(configFolder, profilePath)
	ReadYaml(configPath)
	profileStruct, err := ReadYaml(configPath)
	if err == nil {
		log.Debugf("Configs count: %v\n", len(profileStruct.Profile.Config))
		checkAndSwitch(profileStruct)
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
		// if just 2 profiles are configured switch to the other
		if len(profileStruct.Profile.Config) == 2 {
			if currentIndex == 0 {
				switchTo(profileStruct, 1)
			} else {
				switchTo(profileStruct, 0)
			}
		} else {
			log.Info("Current profile is %v", profileStruct.Profile.Config[currentIndex].Name)
		}
	} else {
		log.Info("None of the configured profiles is in use currently!")
		selectAndSwitch(profileStruct)
	}
}

func selectAndSwitch(profileStruct ProfileStruct) {
	len := len(profileStruct.Profile.Config)
	optionsArray := make([]string, len)
	for i := 0; i < len; i++ {
		optionsArray[i] = profileStruct.Profile.Config[i].Name
	}
	selectedProfileIndex, err :=
		SelectStringCLI("Select profile to switch to", optionsArray)
	if err == nil {
		log.Debugf("Profile selected: %v", selectedProfileIndex)
		switchTo(profileStruct, selectedProfileIndex)
	}
}

func check(profileStruct ProfileStruct) (int, error) {
	currentProfile, err := ReadCurrentProfile(configFolder, currentProfileFile)
	if err == nil {
		for i := 0; i < len(profileStruct.Profile.Config); i++ {
			if currentProfile == profileStruct.Profile.Config[i].Name {
				return i, nil
			}
		}
	}
	// none matched
	return -1, errors.New("no match for profiles")

}

func switchTo(profileStruct ProfileStruct, profileIndex int) {
	// find curr

	for i := 0; i < len(profileStruct.Profile.Config[profileIndex].ConfigFile); i++ {
		typeName := profileStruct.Profile.Config[profileIndex].ConfigFile[i].Type
		fromPath := profileStruct.Profile.Config[profileIndex].ConfigFile[i].FromPath
		fromPath = ResolvePathInConfig(configFolder, fromPath)

		toPath, err := LookupToPath(profileStruct, typeName)
		if err != nil {
			log.Warnf("config-template not found for type %v\n", typeName)
		} else {
			CopyFile(fromPath, toPath)
		}
	}

	// if profileStruct.Profile[profileIndex].GitUserConfig != "" {
	// 	gitResolved = ResolvePathInConfig(configFolder, profileStruct.Profile[profileIndex].GitUserConfig)
	// }
	// CopyFile(gitResolved, "~/.gitconfig")
	// if profileStruct.Profile[profileIndex].SshUserConfig != "" {
	// 	sshResolved = ResolvePathInConfig(configFolder, profileStruct.Profile[profileIndex].SshUserConfig)
	// }
	// CopyFile(sshResolved, "~/.ssh/config")
	profileName := profileStruct.Profile.Config[profileIndex].Name
	log.Info("Switched to profile " + profileName)
	SaveNewProfile(configFolder, currentProfileFile, profileName)
}
