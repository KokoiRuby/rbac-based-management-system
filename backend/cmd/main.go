package main

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/config"
	"github.com/KokoiRuby/rbac-based-management-system/backend/core/logging"
)

func main() {

	err := logging.InitLogger()
	if err != nil {
		panic(err)
	}

	err = config.ParseAndLoad()
	if err != nil {
		panic(err)
	}

	select {}

}
