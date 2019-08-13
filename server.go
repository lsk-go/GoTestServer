package main

import (
	"GoServerTest/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/version/:serverIP/:port", api.Version)
	r.GET("/listcontainers/:serverIP/:port", api.ListContainers)
	r.GET("/listimages/:serverIP/:port", api.Listimages)
	r.Run() // listen and serve on 0.0.0.0:8080
}
