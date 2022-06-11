package routes

import (
	"chat_me_api/src/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(app *gin.Engine) {
	app.POST("/auth/signup", controllers.Signup())
}
