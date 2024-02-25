package main

import (
	"github.com/spf13/viper"

	"github.com/PhilWhittingham/vocab-helper-de/server"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	server.Run(viper.GetString("PORT"))
}
