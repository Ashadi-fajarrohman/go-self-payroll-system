package main

import (
	"go-self-payroll-system/config"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("*env")
	if err != nil {
		log.Fatal(err)
	}

	conf := config.NewConfig()
	server := InitServer(conf)

	go func ()  {
		server.Running()
	}()
}