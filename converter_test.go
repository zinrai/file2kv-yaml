package main

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestConvertPathToYamlName(t *testing.T) {
	tests := []struct {
		name             string
		input            string
		expectedFileName string
		expectedKey      string
	}{
		{
			name:             "simple file path",
			input:            "./hcnfrlw/zorrzzwpgmr.jbk",
			expectedFileName: "hcnfrlw_zorrzzwpgmr_jbk.yaml",
			expectedKey:      "hcnfrlw_zorrzzwpgmr_jbk",
		},
		{
			name:             "path with hyphens",
			input:            "./data/log-file.txt",
			expectedFileName: "data_log_file_txt.yaml",
			expectedKey:      "data_log_file_txt",
		},
		{
			name:             "nested directories",
			input:            "./nbi-bzwnrjr-ihmglbb/ypezqwx/20250107_put-ueqdkr/srlsel_rtkdkdb.pnf",
			expectedFileName: "nbi_bzwnrjr_ihmglbb_ypezqwx_20250107_put_ueqdkr_srlsel_rtkdkdb_pnf.yaml",
			expectedKey:      "nbi_bzwnrjr_ihmglbb_ypezqwx_20250107_put_ueqdkr_srlsel_rtkdkdb_pnf",
		},
		{
			name:             "file without leading ./",
			input:            "file.txt",
			expectedFileName: "file_txt.yaml",
			expectedKey:      "file_txt",
		},
		{
			name:             "file with multiple dots",
			input:            "./config/app.config.json",
			expectedFileName: "config_app_config_json.yaml",
			expectedKey:      "config_app_config_json",
		},
		{
			name:             "hidden file",
			input:            "./.hidden/file.txt",
			expectedFileName: "_hidden_file_txt.yaml",
			expectedKey:      "_hidden_file_txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileName, key := ConvertPathToYamlName(tt.input)
			if fileName != tt.expectedFileName {
				t.Errorf("ConvertPathToYamlName() fileName = %v, want %v", fileName, tt.expectedFileName)
			}
			if key != tt.expectedKey {
				t.Errorf("ConvertPathToYamlName() key = %v, want %v", key, tt.expectedKey)
			}
		})
	}
}

func TestGenerateYamlStructure(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		content string
	}{
		{
			name:    "simple content",
			key:     "test_file",
			content: "Hello, World!",
		},
		{
			name:    "multiline content",
			key:     "config_file",
			content: "line1\nline2\nline3",
		},
		{
			name:    "empty content",
			key:     "empty_file",
			content: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := GenerateYamlStructure(tt.key, tt.content)
			if data.Key != tt.key {
				t.Errorf("GenerateYamlStructure() Key = %v, want %v", data.Key, tt.key)
			}
			if data.Value != tt.content {
				t.Errorf("GenerateYamlStructure() Value = %v, want %v", data.Value, tt.content)
			}
		})
	}
}

func TestMarshalYaml(t *testing.T) {
	tests := []struct {
		name    string
		data    *YamlData
		wantErr bool
	}{
		{
			name: "valid yaml data",
			data: &YamlData{
				Key:   "test_file",
				Value: "test content",
			},
			wantErr: false,
		},
		{
			name: "multiline value",
			data: &YamlData{
				Key:   "multiline_file",
				Value: "line1\nline2\nline3",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yamlBytes, err := MarshalYaml(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalYaml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Verify the YAML can be unmarshaled back
				var result YamlData
				if err := yaml.Unmarshal(yamlBytes, &result); err != nil {
					t.Errorf("Failed to unmarshal generated YAML: %v", err)
				}
				if result.Key != tt.data.Key {
					t.Errorf("Unmarshaled Key = %v, want %v", result.Key, tt.data.Key)
				}
				if result.Value != tt.data.Value {
					t.Errorf("Unmarshaled Value = %v, want %v", result.Value, tt.data.Value)
				}
			}
		})
	}
}
