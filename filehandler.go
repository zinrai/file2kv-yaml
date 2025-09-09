package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

// ReadFileContent reads the content of a file
// Returns an error if the file is binary (not valid UTF-8)
func ReadFileContent(filePath string) (string, error) {
	// Read file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	// Check if content is valid UTF-8
	if !utf8.Valid(content) {
		return "", fmt.Errorf("file is not valid UTF-8 text")
	}

	// Return as string
	return string(content), nil
}

// WriteYamlFile writes YAML data to a file
func WriteYamlFile(outputPath string, data *YamlData) error {
	// Marshal YAML data
	yamlContent, err := MarshalYaml(data)
	if err != nil {
		return fmt.Errorf("failed to marshal YAML: %w", err)
	}

	// Write to file
	if err := os.WriteFile(outputPath, yamlContent, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
