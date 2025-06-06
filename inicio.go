package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hdjwebs/aprender_web/routers"
)

func main() {
	route := gin.Default()

	route.Static("/static", "./static")
	
	route.LoadHTMLGlob("view/*")


	routers.SetupRoutes(route)

	route.Run()
}
