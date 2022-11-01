package main

import (
	"fmt"
	"github.com/spf13/viper"
	"spoti/crawl/server"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	fmt.Println("Hello world!")

	server.Start()
}