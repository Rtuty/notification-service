package main

import (
	"fmt"
	"notification-service/internal/storage"
	"path/filepath"

	"github.com/spf13/viper"
)

func main() {
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

	sc := storage.StorageConfig{
		Host:     viper.GetString("HOST"),
		Port:     viper.GetString("PORT"),
		Database: viper.GetString("DATABASE"),
		Username: viper.GetString("USERNAME"),
		Password: viper.GetString("PASSWORD"),
	}

	fmt.Println(sc)
}
