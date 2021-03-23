package config

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestConfigFromJSONFile(t *testing.T) {
	json_path := "../config.json"
	expected := "{\"host\":\"https://127.0.0.1\", \"port\":80, \"timeout\":30}"

	c, _ := parseConfigFromFile(json_path)
	s, _ := json.Marshal(c)

	assert.JSONEq(t, string(s), expected)
}

func TestConfigFromYAMLFile(t *testing.T) {
	yaml_path := "../config.yaml"
	expected := "{\"host\":\"https://geekbrains.com\", \"port\":8080, \"timeout\":2}"

	c, _ := parseConfigFromFile(yaml_path)
	s, _ := yaml.Marshal(c)

	assert.YAMLEq(t, string(s), expected)
}
