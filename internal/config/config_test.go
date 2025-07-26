package config

import (
	"os"
	"testing"
)

func TestLoad_DefaultValues(t *testing.T) {
	// Clear environment variables
	os.Unsetenv("MITIE_MODEL_PATH")
	os.Unsetenv("PORT")

	config := Load()

	expectedModelPath := "models/ner_model.dat"
	expectedPort := "8080"

	if config.ModelPath != expectedModelPath {
		t.Errorf("Expected ModelPath %s, but got %s", expectedModelPath, config.ModelPath)
	}

	if config.Port != expectedPort {
		t.Errorf("Expected Port %s, but got %s", expectedPort, config.Port)
	}
}

func TestLoad_EnvironmentVariables(t *testing.T) {
	// Set environment variables
	testModelPath := "/custom/path/model.dat"
	testPort := "9090"

	os.Setenv("MITIE_MODEL_PATH", testModelPath)
	os.Setenv("PORT", testPort)

	// Clean up after test
	defer func() {
		os.Unsetenv("MITIE_MODEL_PATH")
		os.Unsetenv("PORT")
	}()

	config := Load()

	if config.ModelPath != testModelPath {
		t.Errorf("Expected ModelPath %s, but got %s", testModelPath, config.ModelPath)
	}

	if config.Port != testPort {
		t.Errorf("Expected Port %s, but got %s", testPort, config.Port)
	}
}

func TestLoad_PartialEnvironmentVariables(t *testing.T) {
	// Set only one environment variable
	testModelPath := "/custom/model.dat"
	expectedPort := "8080" // Default value

	os.Setenv("MITIE_MODEL_PATH", testModelPath)
	os.Unsetenv("PORT")

	// Clean up after test
	defer os.Unsetenv("MITIE_MODEL_PATH")

	config := Load()

	if config.ModelPath != testModelPath {
		t.Errorf("Expected ModelPath %s, but got %s", testModelPath, config.ModelPath)
	}

	if config.Port != expectedPort {
		t.Errorf("Expected Port %s, but got %s", expectedPort, config.Port)
	}
}

func TestLoad_EmptyEnvironmentVariables(t *testing.T) {
	// Set empty environment variables (should use defaults)
	os.Setenv("MITIE_MODEL_PATH", "")
	os.Setenv("PORT", "")

	// Clean up after test
	defer func() {
		os.Unsetenv("MITIE_MODEL_PATH")
		os.Unsetenv("PORT")
	}()

	config := Load()

	expectedModelPath := "models/ner_model.dat"
	expectedPort := "8080"

	if config.ModelPath != expectedModelPath {
		t.Errorf("Expected ModelPath %s, but got %s", expectedModelPath, config.ModelPath)
	}

	if config.Port != expectedPort {
		t.Errorf("Expected Port %s, but got %s", expectedPort, config.Port)
	}
}

func TestConfig_Structure(t *testing.T) {
	config := &Config{
		ModelPath: "/test/path",
		Port:      "3000",
	}

	if config.ModelPath != "/test/path" {
		t.Errorf("Expected ModelPath /test/path, but got %s", config.ModelPath)
	}

	if config.Port != "3000" {
		t.Errorf("Expected Port 3000, but got %s", config.Port)
	}
}