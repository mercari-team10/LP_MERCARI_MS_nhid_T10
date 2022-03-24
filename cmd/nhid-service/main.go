package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/dryairship/techmeet22-nhid-service/config"
	"github.com/dryairship/techmeet22-nhid-service/router"
)

func main() {
	r := gin.Default()
	router.SetUpRoutes(r)

	if err := r.Run(":" + config.ApplicationPort); err != nil {
		log.Fatalln("[ERROR] Could not start the server: ", err.Error())
	}
}
