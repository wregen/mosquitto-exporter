package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type MosquittoAuthentication struct {
	SuperUserLogin    string `json:"superUserLogin,omitempty"`
	SuperUserPassword string `json:"superUserPassword,omitempty"`
}

type MqttConfig struct {
	MosquittoAuthentication MosquittoAuthentication
}

func getConfig(fileName string) (user string, pass string, err error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return "", "", err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return "", "", err

	}

	var mqttConfig MqttConfig
	err = json.Unmarshal(byteValue, &mqttConfig)
	if err != nil {
		return "", "", err
	}

	return mqttConfig.MosquittoAuthentication.SuperUserLogin, mqttConfig.MosquittoAuthentication.SuperUserPassword, nil
}
