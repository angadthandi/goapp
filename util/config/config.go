package config

import (
	"encoding/json"
	"os"
	"strconv"

	log "github.com/angadthandi/goapp/log"
)

const ConfigFilePath = "config.json"

type ConfigStruct struct {
	SQLiteDBPath  string
	ServerPort    string
	JWTAuthSecret string
	LogPath       string
	LogLevel      string
	LogMaxSize    int
	LogMaxBackups int
	LogMaxAge     int
	LogCompress   bool
}

type configJSONStruct struct {
	SQLiteDBPath  string `json:"sqlitedbpath"`
	ServerPort    string `json:"port"`
	JWTAuthSecret string `json:"jwtauthsecret"`
	LogPath       string `json:"logpath"`
	LogLevel      string `json:"loglevel"`
	LogMaxSize    string `json:"logmaxsize"`
	LogMaxBackups string `json:"logmaxbackups"`
	LogMaxAge     string `json:"logmaxage"`
	LogCompress   string `json:"logcompress"`
}

func Read() (ConfigStruct, error) {
	var (
		data     ConfigStruct
		jsonData configJSONStruct
	)

	file, err := os.Open(ConfigFilePath)
	defer file.Close()
	if err != nil {
		log.Errorf("unable to open config file: %v", err)
		return data, err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&jsonData)
	if err != nil {
		log.Errorf("unable to decode config: %v", err)
		return data, err
	}

	data.SQLiteDBPath = jsonData.SQLiteDBPath
	data.ServerPort = jsonData.ServerPort
	data.JWTAuthSecret = jsonData.JWTAuthSecret
	data.LogPath = jsonData.LogPath
	data.LogLevel = jsonData.LogLevel

	data.LogMaxSize, err = strconv.Atoi(jsonData.LogMaxSize)
	if err != nil {
		log.Errorf("unable to read LogMaxSize: %v", err)
		return data, err
	}
	data.LogMaxBackups, err = strconv.Atoi(jsonData.LogMaxBackups)
	if err != nil {
		log.Errorf("unable to read LogMaxBackups: %v", err)
		return data, err
	}
	data.LogMaxAge, err = strconv.Atoi(jsonData.LogMaxAge)
	if err != nil {
		log.Errorf("unable to read LogMaxAge: %v", err)
		return data, err
	}

	data.LogCompress = true
	if jsonData.LogCompress == "true" {
		data.LogCompress = false
	}

	return data, err
}
