package main

import "github.com/jparrill/gormsample/internal/cmd"

func main() {

	logger := cmd.InitLogger()
	logger.Info("Logging Application from Main")

}
