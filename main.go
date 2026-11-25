package main

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/victorsantosbrazil/go-api-echo-template/docs/swagger"
	app "github.com/victorsantosbrazil/go-api-echo-template/internal"
	"github.com/victorsantosbrazil/go-api-echo-template/pkg/app/log"
)

// @title Accounts API Fork
// @version 0.1.0
// @description API for managing bank accounts and credit cards
func main() {
	logger := log.NewLogger()

	a, err := app.NewApp()
	if err != nil {
		logger.Fatal(err.Error())
	}

	err = a.Start()
	if err != nil {
		logger.Fatal(err.Error())
	}
}
