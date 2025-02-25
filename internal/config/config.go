package config

import (
)

type Config struct {
	Db string `json:"db_url"`
	User string `json:"current_user_name"` 
}

func (c Config) SetUser(name string) error {
	c.User = name
	
	// Update config file with user name change
	err := write(c)
	if (err != nil) {
		return err
	}

	return nil
}
