package config

import (
	"go-self-payroll-system/config/postgres"
	"os"

	"gorm.io/gorm"
)

type (
	config struct{}

	Config interface {
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
