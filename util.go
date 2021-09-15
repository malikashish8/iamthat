package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
)

func ResolvePath(somepath string) string {
	trimmed := strings.TrimSpace(somepath)
	absolutePath := trimmed
	if strings.HasPrefix(trimmed, "~/") {
		localOnly := trimmed[2:]
		userHomeDir, _ := os.UserHomeDir()
		absolutePath = filepath.Join(userHomeDir, localOnly)
	} else if !strings.HasPrefix(trimmed, "/") {
		wd, err := os.Getwd()
		if err == nil {
			absolutePath = filepath.Join(wd, trimmed)
		}
	}
	return absolutePath
}

func CheckFilesSameContent(path1 string, path2 string) bool {
	p1 := ResolvePath(path1)
	p2 := ResolvePath(path2)
	data1, err1 := os.ReadFile(p1)
	data2, err2 := os.ReadFile(p2)
	if err1 != nil && err2 != nil {
		return false
	} else if string(data1) != string(data2) {
		return false
	}
	return true
}

func SelectStringCLI(label string, options []string) (int, error) {
	prompt := promptui.Select{
		Label: label,
		Items: options,
	}

	index, _, err := prompt.Run()
	if err != nil {
		return -1, err
	}
	return index, nil
}

// Copy source file to destination.
// If source path is empty string then delete destination.
func CopyFile(src string, dest string) {
	srcPath := ResolvePath(src)
	destPath := ResolvePath(dest)
	if src != "" {
		bytesRead, err := ioutil.ReadFile(srcPath)
		if err == nil {
			errWrite := ioutil.WriteFile(destPath, bytesRead, 0600)
			if errWrite != nil {
				log.Error("Error writing file " + destPath)
			}
		} else {
			log.Error("Error reading file " + srcPath)
		}
	} else {
		err := os.Remove(destPath)
		if err != nil {
			log.Error("Error deleting file " + destPath)
			log.Error(err)
		}
	}
}

func ResolvePathInConfig(configFolder string, relativePath string) string {
	absolutePath := filepath.Join(ResolvePath(configFolder), relativePath)
	return absolutePath
}

// Get the name of Profile currently used from `currentProfileFile`
// If the file is not found then error is returned
func ReadCurrentProfile(configFolder string, currentProfileFile string) (string, error) {
	profilePath := filepath.Join(ResolvePath(configFolder), currentProfileFile)
	data, err := os.ReadFile(profilePath)
	if err == nil {
		log.Debug("Current Profile: " + string(data))
		return string(data), nil
	}
	return "", err
}

// Save `profileName` in `currentProfileFile` in `configFolder`
// If config folder does not exist it is created
func SaveNewProfile(configFolder string, currentProfileFile string, profileName string) error {
	configFolder = ResolvePath(configFolder)
	profilePath := filepath.Join(configFolder, currentProfileFile)
	// check if config folder exits
	if _, err := os.Stat(configFolder); os.IsNotExist(err) {
		err := os.MkdirAll(configFolder, 0711)
		if err == nil {
			return err
		}
	}
	errWrite := ioutil.WriteFile(profilePath, []byte(profileName), 0600)
	return errWrite
}

// Look up to-path for a given config-file type
func LookupToPath(profileStruct ProfileStruct, typeName string) (string, error) {
	for i := 0; i < len(profileStruct.Profile.ConfigTemplate); i++ {
		if typeName == profileStruct.Profile.ConfigTemplate[i].Name {
			return profileStruct.Profile.ConfigTemplate[i].ToPath, nil
		}
	}
	return "", errors.New("to-path not found for " + typeName)
}
