package storage

import (
	"fmt"
	"path/filepath"

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
	cfg := viper.New()
	cfg.SetConfigName("config")
	cfg.SetConfigType("env")

	root, err := filepath.Abs(".")
	if err != nil {
		panic(fmt.Errorf("ошибка при получении корневой директории: %s", err))
	}

	cfg.AddConfigPath(root)

	err = cfg.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found")
		} else {
			fmt.Printf("Fatal error config file: %s \n", err)
		}
	}

	sc := StorageConfig{
		Host:     cfg.GetString("HOST"),
		Port:     cfg.GetString("PORT"),
		Database: cfg.GetString("DATABASE"),
		Username: cfg.GetString("USERNAME"),
		Password: cfg.GetString("PASSWORD"),
	}

	return sc, nil
}
