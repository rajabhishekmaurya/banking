package main

import (
	"github.com/rajabhishekmaurya/banking/app"
	"github.com/rajabhishekmaurya/banking/logger"
)

func main() {
	logger.Info("Starting the application....")
	app.Start()
}



