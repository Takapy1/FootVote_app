package routers

import (
	"github.com/gin-gonic/gin"
	// "gopkg.in/gin-gonic/gin.v1"

	"github.com/Takapy1/FootVote/server/controllers"
)

func SetUp() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		// admin := api.Group("/admin")
		// {
		// 	games := 
		// }

		games := api.Group("/games")
		{
			games.GET("/", controllers.GetGames)
			games.GET("/:id", controllers.GetGame)
		}
		// users := api.Group("/users")
		// {
		// 	users.GET("/:id", GetUser)
		// }
	}
	return router
}