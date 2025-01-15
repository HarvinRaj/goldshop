package configs

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func createMockConfigFile(t *testing.T, content string) string {
	t.Helper()

	tempFile, err := os.CreateTemp("", "testconfig-*.json")
	if err != nil {
		t.Errorf("Failed to create temp config file, %s", err)
	}

	_, err = tempFile.WriteString(content)
	if err != nil {
		t.Errorf("Failed to write content in temp file %s", err)
	}

	if err = tempFile.Close(); err != nil {
		t.Errorf("Failed to close temp file, %s", err)
	}

	return tempFile.Name()
}

func TestLoadConfigSuccess(t *testing.T) {

	mockConfigFile := `{
		"OSTYPE": "test",
		"PORT": "0000",
		"DBNAME": "TESTER",
		"PARSETIME": false
	}`

	tempFile := createMockConfigFile(t, mockConfigFile)
	defer os.Remove(tempFile)

	expectedConfig := &Config{
		OSType:    "test",
		Port:      "0000",
		DBName:    "TESTER",
		ParseTime: false,
	}

	config, err := LoadConfig(tempFile)
	assert.NotNil(t, config)
	assert.NoError(t, err)
	assert.Equal(t, expectedConfig, config)

}

func TestLoadConfigFileNotFound(t *testing.T) {

	_, err := LoadConfig("non_existent_file.json")
	assert.Error(t, err)
}

func TestLoadConfigInvalidJSON(t *testing.T) {

	//the last comma makes it invalid
	mockConfigFile := `{
		"OSTYPE": "test",
		"PORT": "0000",
		"DBNAME": "TESTER",
		"PARSETIME": false,
	}`

	tempFile := createMockConfigFile(t, mockConfigFile)
	defer os.Remove(tempFile)

	_, err := LoadConfig(tempFile)
	assert.NotNil(t, err)
	assert.Error(t, err)
}
