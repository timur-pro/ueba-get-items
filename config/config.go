package config

import (
	"encoding/json"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"sync"
)

type Config struct {
	// Адрес нашего сервера
	ServerAddress string `envconfig:"LISTEN" default:":8080"`
	// Путь до файла ueba.csv
	FilePath string `envconfig:"UEBA_FILE_PATH" default:"./ueba.csv"`
	// Индекс
	IdColumnIndex int `envconfig:"ID_COLUMN_INDEX" default:"1"`
}

var (
	config Config
	once   sync.Once
)

// GetConfig возвращает объект с конфигурацией
func GetConfig() *Config {
	once.Do(func() {
		err := envconfig.Process("", &config)
		if err != nil {
			log.Fatal(err)
		}

		b, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Configuration:", string(b))
	})

	return &config
}
