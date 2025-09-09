package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Check command line arguments
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <output_directory>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: find ./source -type f | %s ./output\n", os.Args[0])
		os.Exit(1)
	}

	outputDir := os.Args[1]

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to create output directory: %v\n", err)
		os.Exit(1)
	}

	// Get absolute path of output directory for filtering
	outputAbsPath, err := filepath.Abs(outputDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to resolve output directory path: %v\n", err)
		os.Exit(1)
	}

	// Read file paths from stdin
	scanner := bufio.NewScanner(os.Stdin)
	processedCount := 0
	skippedCount := 0

	for scanner.Scan() {
		filePath := strings.TrimSpace(scanner.Text())
		if filePath == "" {
			continue
		}

		// Get absolute path of input file
		fileAbsPath, err := filepath.Abs(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Cannot resolve path %s: %v\n", filePath, err)
			os.Exit(1)
		}

		// Skip files in output directory (this is expected behavior, not an error)
		if strings.HasPrefix(fileAbsPath, outputAbsPath+string(os.PathSeparator)) {
			fmt.Printf("Skipping: %s (in output directory)\n", filePath)
			skippedCount++
			continue
		}

		// Check if file exists
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Error: File not found: %s\n", filePath)
			os.Exit(1)
		}

		// Generate YAML filename
		yamlFileName, key := ConvertPathToYamlName(filePath)
		outputPath := filepath.Join(outputDir, yamlFileName)

		// Read file content
		content, err := ReadFileContent(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Failed to read file %s: %v\n", filePath, err)
			os.Exit(1)
		}

		// Generate YAML structure
		yamlData := GenerateYamlStructure(key, content)

		// Write YAML file
		if err := WriteYamlFile(outputPath, yamlData); err != nil {
			fmt.Fprintf(os.Stderr, "Error: Failed to write YAML file %s: %v\n", outputPath, err)
			os.Exit(1)
		}

		fmt.Printf("Processed: %s -> %s\n", filePath, yamlFileName)
		processedCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to read from stdin: %v\n", err)
		os.Exit(1)
	}

	// Print summary
	fmt.Printf("\nCompleted: %d files processed successfully\n", processedCount)
	if skippedCount > 0 {
		fmt.Printf("Skipped: %d files (in output directory)\n", skippedCount)
	}
}
