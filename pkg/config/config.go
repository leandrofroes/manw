package config

import (
	"os"
	"io/ioutil"

	"manw/pkg/utils"
	"gopkg.in/yaml.v2"
)

type Config struct{
	CacheDir	string	`yaml:"Cache Directory"`
}

func Load() (cachePath string){
	configFile := "manw.yml"

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		utils.Warning("Unable to find the config file.")
	}

	var config Config
	source, err := ioutil.ReadFile(configFile)
	utils.CheckError(err)
	err = yaml.Unmarshal(source, &config)
	utils.CheckError(err)

	cachePath = config.CacheDir

	if _, err := os.Stat(cachePath); os.IsNotExist(err) {
		err := os.Mkdir(cachePath, 0755)
		utils.CheckError(err)
	}

	return cachePath
}
