// config.go
package configuration

import (
	"encoding/json"
	"os"
)

// Config struct represents the structure of your configuration file
type Config struct {
	Authorization struct {
		AuthorizedUsername string `json:"authorizedUsername"`
		AuthorizedPassword string `json:"authorizedpassword"`
	} `json:"authorization"`
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Dbname   string `json:"dbname"`
	} `json:"database"`
}

// LoadConfig reads the configuration from a JSON file
func LoadConfig(filename string) (Config, error) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
