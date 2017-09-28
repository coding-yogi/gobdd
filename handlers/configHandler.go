package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.wdf.sap.corp/ml-base/lr-bdd-tests/models/appconfig"
)

func getConfigJSON() appconfig.AppConfig {
	file, err := ioutil.ReadFile("../config.json")
	if err != nil {
		log.Fatal("Unable to read config file")
	}

	appConf := appconfig.AppConfig{}
	json.Unmarshal(file, &appConf)

	return appConf
}

//GetEnvDetails ...
func GetEnvDetails(envName string) (appconfig.Environment, error) {
	appConf := getConfigJSON()
	for _, env := range appConf.Envs {
		if env.Name == envName {
			return env, nil
		}
	}

	return appconfig.Environment{}, errors.New("Environment not found")
}
