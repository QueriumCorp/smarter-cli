/*
Copyright © 2024 Lawrence McDaniel lawrence@querium.com
*/
package main

import (
	"log"

	"github.com/QueriumCorp/smarter-cli/cmd"
	_ "github.com/QueriumCorp/smarter-cli/cmd/delete"
	_ "github.com/QueriumCorp/smarter-cli/cmd/delete/account"
	_ "github.com/QueriumCorp/smarter-cli/cmd/get"
	_ "github.com/QueriumCorp/smarter-cli/cmd/get/account"
	_ "github.com/QueriumCorp/smarter-cli/cmd/manifest"
	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cmd.Execute()

	// Set up a global --environment flag and bind this to viper.
	pflag.String("environment", "prod", "API environment: local, alpha, beta, next, prod")
	if err := viper.BindPFlag("environment", pflag.Lookup("environment")); err != nil {
		log.Fatalf("Error binding flag: %v", err)
	}
}
