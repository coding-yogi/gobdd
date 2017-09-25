package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.com/coding-yogi/go_bdd/models"
)

func getConfigJSON() models.AppConfig {
	file, err := ioutil.ReadFile("../config.json")
	if err != nil {
		log.Fatal("Unable to read config file")
	}

	appConf := models.AppConfig{}
	json.Unmarshal(file, &appConf)

	return appConf
}

//GetEnvDetails ...
func GetEnvDetails(envName string) (models.Environment, error) {
	appConf := getConfigJSON()
	for _, env := range appConf.Envs {
		if env.Name == envName {
			return env, nil
		}
	}

	return models.Environment{}, errors.New("Environment not found")
}
