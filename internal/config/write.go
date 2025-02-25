package config

import (
	"os"
	"fmt"
	"encoding/json"
)

func write(cfg Config) error {
	jsonData, err := json.Marshal(cfg)
	if (err != nil) {
		fmt.Printf("Failed to marhsal data.\nErr: %v", err)
		return err
	}

	path, err := getConfigFilePath()
	if (err != nil) {
		fmt.Printf("Failed to get config path.\nErr: %v", err)
		return err
	}

	if err = os.WriteFile(path, jsonData, 0644); err != nil {
		fmt.Printf("Failed to write to config.\nErr: %v", err)
		return err
	}

	return nil
}
