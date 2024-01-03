package main

import (
	"github.com/dreking/event-booking-api/api/routes"
	"github.com/dreking/event-booking-api/api/utils"
	"github.com/dreking/event-booking-api/db"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(config.ServerAddress)
}
