package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadFileContent(t *testing.T) {
	// Create temporary test files
	tempDir := t.TempDir()

	tests := []struct {
		name        string
		fileName    string
		content     []byte
		wantErr     bool
		errContains string
	}{
		{
			name:     "valid UTF-8 text file",
			fileName: "valid.txt",
			content:  []byte("Hello, World!\nThis is a test."),
			wantErr:  false,
		},
		{
			name:     "empty file",
			fileName: "empty.txt",
			content:  []byte(""),
			wantErr:  false,
		},
		{
			name:        "invalid UTF-8",
			fileName:    "invalid.dat",
			content:     []byte{0xFF, 0xFE, 0x00, 0x00, 0xFF, 0xFF},
			wantErr:     true,
			errContains: "not valid UTF-8",
		},
		{
			name:     "UTF-8 with special characters",
			fileName: "unicode.txt",
			content:  []byte("こんにちは\n世界\n🌍"),
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test file
			filePath := filepath.Join(tempDir, tt.fileName)
			if err := os.WriteFile(filePath, tt.content, 0644); err != nil {
				t.Fatalf("Failed to create test file: %v", err)
			}

			// Test ReadFileContent
			result, err := ReadFileContent(filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFileContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				if tt.errContains != "" && err != nil {
					if !contains(err.Error(), tt.errContains) {
						t.Errorf("ReadFileContent() error = %v, should contain %v", err, tt.errContains)
					}
				}
			} else {
				if result != string(tt.content) {
					t.Errorf("ReadFileContent() = %v, want %v", result, string(tt.content))
				}
			}
		})
	}

	// Test non-existent file
	t.Run("non-existent file", func(t *testing.T) {
		_, err := ReadFileContent("/non/existent/file.txt")
		if err == nil {
			t.Errorf("ReadFileContent() should return error for non-existent file")
		}
	})
}

func TestWriteYamlFile(t *testing.T) {
	tempDir := t.TempDir()

	tests := []struct {
		name    string
		data    *YamlData
		wantErr bool
	}{
		{
			name: "simple yaml",
			data: &YamlData{
				Key:   "test_file",
				Value: "test content",
			},
			wantErr: false,
		},
		{
			name: "multiline yaml",
			data: &YamlData{
				Key:   "multiline",
				Value: "line1\nline2\nline3",
			},
			wantErr: false,
		},
		{
			name: "empty value",
			data: &YamlData{
				Key:   "empty",
				Value: "",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outputPath := filepath.Join(tempDir, tt.data.Key+".yaml")

			err := WriteYamlFile(outputPath, tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteYamlFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Verify file exists and content is correct
				content, err := os.ReadFile(outputPath)
				if err != nil {
					t.Errorf("Failed to read written file: %v", err)
					return
				}

				// Check that it's valid YAML
				if len(content) == 0 {
					t.Errorf("Written file is empty")
				}

				// Verify the content contains the key
				if !contains(string(content), "key: "+tt.data.Key) {
					t.Errorf("Written YAML doesn't contain expected key: %s", tt.data.Key)
				}
			}
		})
	}
}

// Helper function
func contains(s, substr string) bool {
	return len(substr) > 0 && len(s) >= len(substr) &&
		(s == substr || len(s) > len(substr) &&
			(s[:len(substr)] == substr || s[len(s)-len(substr):] == substr ||
				len(s) > len(substr) && containsMiddle(s, substr)))
}

func containsMiddle(s, substr string) bool {
	for i := 1; i < len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
