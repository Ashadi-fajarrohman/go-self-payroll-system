package config

import (
	"go-self-payroll-system/config/postgres"
	"os"
	"strconv"

	"gorm.io/gorm"
)

type (
	config struct{}

	Config interface {
		Database() *gorm.DB
		ServiceName() string
		ServicePort() int
	}
)

func NewConfig() Config {
	return &config{}
}

func (c *config) Database() *gorm.DB {
	return postgres.InitGorm()
}

func (c *config) ServiceName() string {
	return os.Getenv("SERVICE_NAME")
}

func (C *config) ServicePort() int {
	port := os.Getenv("SERVICE_PORT")
	portInt, _ := strconv.Atoi(port)

	return portInt
}