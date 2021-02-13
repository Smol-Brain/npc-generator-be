package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"npc-generator-be/global"
	"npc-generator-be/npc"
	"npc-generator-be/user"
)

// InitializeRoutes sets up routing for server requests
func InitializeRoutes() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, `
			List of endpoints:
			
			GET /users - Retrieve all users
			GET /users/:id - Retrieve specific user
			POST /users - Create a user

			GET /npcs - Retrieve all npcs
			GET /npcs/:id - Retrieve specific npc
			POST /npcs - Create an npc (optional int req.body.npc_amount generates x npcs)
		`)
	})

	npc.Route(router)
	user.Route(router)

	router.Run(":" + global.Conf.Port)
}
