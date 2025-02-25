package main

import (
	"fmt"
	"os"

	"github.com/AhmetMuratAcar/blog-agg/internal/config"
)

func main() {
	userConfig, err := config.Read()
	if (err != nil) {
		os.Exit(1)
	}

	fmt.Printf("---START---\ndb: %s\nuser: %s\n", userConfig.Db, userConfig.User)

	err = userConfig.SetUser("Joe")
	if (err != nil) {
		os.Exit(1)
	}

	userConfig, err = config.Read()
	if (err != nil) {
		os.Exit(1)
	}

	fmt.Printf("---END---\ndb: %s\nuser: %s\n", userConfig.Db, userConfig.User)
}
