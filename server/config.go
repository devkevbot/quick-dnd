package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config holds important configuration data, such as database
// credentials. In the future, this type may be expanded to hold server
// configuration data as well.
type Config struct {
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
		DBName   string `yaml:"dbname"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	} `yaml:"database"`
	IsProduction bool `yaml:"production"`
	HTTPServer   struct {
		Port int `yaml:"port"`
	} `yaml:"http_server"`
	JWTSigningKey string `yaml:"jwt_signing_key"`
}

// CreatePostgreSQLDBConnString returns a formatted string used the
// database present in `c`. Uses `sslmode` is `useSSL` is set to true.
func (c Config) CreatePostgreSQLDBConnString(useSSL bool) string {
	sslmode := "disable"
	if useSSL {
		sslmode = "enable"
	}

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.Username,
		c.Database.Password,
		c.Database.DBName,
		sslmode,
	)
}

func createConfigFromFile(path string) *Config {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	c := Config{}
	if err := yaml.NewDecoder(file).Decode(&c); err != nil {
		log.Fatal(err)
	}
	return &c
}
