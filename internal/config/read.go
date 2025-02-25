package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if (err != nil) {
		fmt.Printf("Failed getting location of home dir.\nErr: %v\n", err)
		return Config{}, err
	}

	jsonFile, err := os.Open(path)
	if (err != nil) {
		fmt.Printf("Failed opening JSON file.\nErr: %v", err)
		return Config{}, err
	}
	defer jsonFile.Close()
	
	jsonData, err := io.ReadAll(jsonFile)
	if (err != nil) {
		fmt.Printf("Failed reading JSON file.\nErr: %v", err)
		return Config{}, err
	}

	var userConfig Config
	if err := json.Unmarshal(jsonData, &userConfig); err != nil {
		fmt.Printf("Failed unmarshalling json data.\nErr:%v\n", err)
		return Config{}, err
	}

	return userConfig, nil
}
