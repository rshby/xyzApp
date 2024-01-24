package config

import (
	"github.com/stretchr/testify/mock"
	"xyzApp/app/config"
)

type ConfigMock struct {
	Mock mock.Mock
}

func NewConfigMock() *ConfigMock {
	return &ConfigMock{Mock: mock.Mock{}}
}

func (c *ConfigMock) GetConfig() *config.AppConfig {
	args := c.Mock.Called()

	value := args.Get(0)
	if value == nil {
		return nil
	}

	return value.(*config.AppConfig)
}
