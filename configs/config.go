package configs

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	SMS       string `json:"SMS"`
	VoiceCall string `json:"voice_call"`
	Billing   string `json:"billing"`
	MMS       string `json:"MMS"`
	Support   string `json:"support"`
	Incident  string `json:"incident"`
	Email     string `json:"email"`
}

func GetConfig() Config {
	dir, err := filepath.Abs("")
	if err != nil {
		log.Printf("Ошибка пути к файлу настроек:\n%v", err)
	}
	pathConfigFile := filepath.Join(dir, "configs", "config.json")
	log.Println(pathConfigFile)
	configByte, err := os.ReadFile(pathConfigFile)
	if err != nil {
		log.Printf("Ошибка получения настроек:\n%v", err)
		os.Exit(1)
	}
	var config Config
	if err = json.Unmarshal(configByte, &config); err != nil {
		log.Printf("Ошибка получения настроек:\n%v", err)
		os.Exit(1)
	}
	return config
}