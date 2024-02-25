package main

import (
	"github.com/PhilWhittingham/vocab-helper-de/server"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	server.Run(viper.GetString("PORT"))
}
