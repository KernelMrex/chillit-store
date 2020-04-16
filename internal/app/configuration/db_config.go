package configuration

import (
	"fmt"
)

type DBConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Database string `yaml:"db"`
}

func (c DBConfig) GetUrl() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)
}
