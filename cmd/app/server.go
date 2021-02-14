package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"npc-generator-be/cmd/npc"
	"npc-generator-be/cmd/user"
	"npc-generator-be/internal/config"
)

// InitializeServer sets up engine and routing for server requests
func InitializeServer(config config.Config, db *gorm.DB) (router *gin.Engine) {
	router = gin.Default()

	// Attach db reference for database operations
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Defaulting root url to list of commands for now
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

	// Attach resource-specific routes
	npc.Route(router)
	user.Route(router)

	err := router.Run(":" + config.PORT)
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}

	log.Printf("Server running on port %v", config.PORT)

	return
}
