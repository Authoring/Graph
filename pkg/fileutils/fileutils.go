package fileutils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"

	"github.com/Authoring/Graph/logger"
)

var baseFolder = "~/.graph"

// GetConfigPath returns the configuration path
func GetConfigPath(name string) string {
	initBaseFolder()
	return path.Join(baseFolder, name)
}

// WriteFile writes to a file
func WriteFile(base, name string, data []byte) error {
	initBaseFolder()
	_, filepath := GetBaseAndFilePath(base, name)
	return ioutil.WriteFile(filepath, data, os.ModePerm)
}

// GetBaseAndFilePath get the base and filepath
func GetBaseAndFilePath(base string, name string) (string, string) {
	p := path.Join(baseFolder, fmt.Sprintf("%s.db", base))
	filepath := path.Join(p, fmt.Sprintf("%s.yml", name))

	return p, filepath
}

// LoadAndCreate loads a file by file name
func LoadAndCreate(base, name string) ([]byte, error) {
	initBaseFolder()

	if !FileExists(baseFolder) {
		createBaseFolder()
	}

	p, filepath := GetBaseAndFilePath(base, name)

	if !FileExists(p) {
		err := os.MkdirAll(p, os.ModePerm)

		if err != nil {
			return nil, err
		}
	}

	if !FileExists(filepath) {
		err := ioutil.WriteFile(filepath, nil, os.ModePerm)

		if err != nil {
			return nil, err
		}
	}

	return ioutil.ReadFile(filepath)
}

// initBaseFolder sets the current folder
func initBaseFolder() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	baseFolder = path.Join(usr.HomeDir, ".graph")
}

// createBaseFolder creates the base folder
func createBaseFolder() {
	err := os.MkdirAll(baseFolder, os.ModePerm)
	if err != nil {
		logger.L.Panicf(err.Error())
	}
}

// FileExists verify if the main folder exists
func FileExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}

	return false
}
