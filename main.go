package main

import (
	"RGT/konis/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/img/profile", "./img/profile")
	r.Static("/img/product", "./img/product")

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	routers.RouterCombain(r)

	r.Run("0.0.0.0:8000")

}
