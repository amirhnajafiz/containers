package pkg

import (
	"encoding/json"
	"os"
)

// defaultConfigs returns a map of default configurations for the container.
func defaultConfigs() map[string]string {
	return map[string]string{
		"memory": "128",
		"cpu":    "512",
	}
}

// ReadConfigs reads the configuration file and returns a map of configurations.
func ReadConfigs() (map[string]string, error) {
	configs := defaultConfigs()

	// read the configuration file
	file, err := os.Open("configs.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// decode the JSON into the configs map
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&configs); err != nil {
		return nil, err
	}

	return configs, nil
}
