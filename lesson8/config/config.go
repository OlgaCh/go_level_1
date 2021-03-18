package config

import (
	"encoding/json"
	"errors"
	"flag"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
)

func getEnv(key string, defaultVal string) string {
	if envVal, ok := os.LookupEnv(key); ok {
		return envVal
	}
	return defaultVal
}

func getEnvInt64(key string, defaultVal int64) int64 {
	if envVal, ok := os.LookupEnv(key); ok {
		envInt64, err := strconv.ParseInt(envVal, 10, 64)
		if err == nil {
			return envInt64
		}
	}
	return defaultVal
}

type Config struct {
	Host    string `yaml:"host" json:"host"`
	Port    int64  `yaml:"port" json:"port"`
	Timeout int64  `yaml:"timeout" json:"timeout"`
}

func New() (*Config, error) {
	var (
		host    = flag.String("host", getEnv("HOST", "https://127.0.0.1"), "URL of the service")
		port    = flag.Int64("port", getEnvInt64("PORT", 8080), "Port")
		timeout = flag.Int64("timeout", getEnvInt64("TIMEOUT", 120), "Connection Timeout")
		file    = flag.String("file", "", "Path to yaml or json config file")
	)
	flag.Parse()

	var conf *Config
	var err error

	if *file != "" {
		log.Println("Getting config information from file.")
		conf, err = parseConfigFromFile(*file)
		if err != nil {
			log.Fatal("Error getting config from file")
		}
	} else {
		conf = &Config{
			Host:    *host,
			Port:    *port,
			Timeout: *timeout,
		}
	}

	isValid := conf.Validate()
	if isValid {
		return conf, nil
	}
	return conf, errors.New("Malformed Config")
}

func parseConfigFromFile(filePath string) (*Config, error) {
	var c Config

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return &c, err
	}

	switch strings.ToLower(path.Ext(filePath)) {
	case "yaml, yml":
		err = yaml.Unmarshal(b, &c)
	case "json":
		err = json.Unmarshal(b, &c)
	}

	return &c, err
}

func (c Config) Validate() bool {
	if c.Timeout == 0 {
		log.Fatal("Timeout should be greater than zero")
		return false
	}
	if c.Port != 80 && c.Port != 8080 {
		log.Fatal("Acceptable ports are 80 and 8080. Provided:", c.Port)
		return false
	}
	_, err := url.ParseRequestURI(c.Host)
	if err != nil {
		log.Fatal("Invalid host:", c.Host)
		return false
	}
	return true
}
