package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hdjwebs/aprender_web/controller"
)

func SetupRoutes(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{
			"titulo":  "Inicio",
			"mensage": "Bienvenido a mi sitio web",
		})
	})

	router.POST("/login", func(c *gin.Context) {
		controller.Login(c)
	})

}
