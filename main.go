package main

import (
	"github.com/e-commerce-api/config"
	"github.com/e-commerce-api/router"
)

func init() {
	config.InitiliazeConfig()
	config.InitDB()
	config.SyncDB()
}
func main() {
	router.Server().Run()
}
