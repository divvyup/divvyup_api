package main

import (
	"fmt"

	"github.com/domtheporcupine/divvyup_api/config"
)

func main() {
	config.AppConfig()
	fmt.Println(config.Secret())
}
