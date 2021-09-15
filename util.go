package main

import (
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
				log.Error("Error writing file " + srcPath)
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
