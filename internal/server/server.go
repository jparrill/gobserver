package server

import "github.com/jparrill/gobserver/internal/config"

func Init() {
	router := NewRouter()
	router.Run(config.CFG.AppPort)
}
