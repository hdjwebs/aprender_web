package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	email := c.PostForm("email")
	senha := c.PostForm("senha")

	if email == "admin" && senha == "1234" {
		c.HTML(200, "index.html", gin.H{
			"email":   email,
			"titulo":  "Inicio",
			"mensage": "Bienvenido a mi sitio web",
		})
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"Erro": "Usuário ou senha inválidos.",
		})
	}
}
