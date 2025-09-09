package main

import (
	"strings"

	"gopkg.in/yaml.v3"
)

// YamlData represents the structure of the YAML file
type YamlData struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

// ConvertPathToYamlName converts a file path to a YAML filename and key
// Returns: (YAML filename, key name)
func ConvertPathToYamlName(filePath string) (string, string) {
	// Remove leading "./"
	cleanPath := strings.TrimPrefix(filePath, "./")

	// Apply conversion rules:
	// 1. Replace "/" with "_"
	// 2. Replace "-" with "_"
	// 3. Replace "." with "_"
	converted := cleanPath
	converted = strings.ReplaceAll(converted, "/", "_")
	converted = strings.ReplaceAll(converted, "-", "_")
	converted = strings.ReplaceAll(converted, ".", "_")

	// Generate YAML filename and key
	yamlFileName := converted + ".yaml"
	key := converted

	return yamlFileName, key
}

// GenerateYamlStructure creates a YAML data structure
func GenerateYamlStructure(key, content string) *YamlData {
	return &YamlData{
		Key:   key,
		Value: content,
	}
}

// MarshalYaml converts YamlData struct to YAML format
func MarshalYaml(data *YamlData) ([]byte, error) {
	return yaml.Marshal(data)
}
