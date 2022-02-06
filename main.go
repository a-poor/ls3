package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var BucketName string

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	BucketName = os.Getenv("BUCKET_NAME")
}

func main() {
	// Set up the CLI config data
	viper.SetConfigName(".ls3.yaml")
	viper.AddConfigPath("$HOME")

	viper.SetDefault("Message", "Hello, World!")

	// Create the app
	app := newApp()

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
