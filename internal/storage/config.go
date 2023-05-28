package storage

import (
	"fmt"

	"github.com/spf13/viper"
)

type StorageConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetStorageConfig() (StorageConfig, error) {
	viper.SetConfigFile("config.env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("не удалось прочитать файл конфигурации: %s", err))
	}

	sc := StorageConfig{
		Host:     viper.GetString("HOST"),
		Port:     viper.GetString("PORT"),
		Database: viper.GetString("DATABASE"),
		Username: viper.GetString("USERNAME"),
		Password: viper.GetString("PASSWORD"),
	}

	return sc, nil
}
