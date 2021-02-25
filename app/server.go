package app

import (
	"github/Smol-Brain/npc-generator-be/npc"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// InitializeServer sets up engine and routing for server requests
func InitializeServer(config Config, db *gorm.DB) (router *gin.Engine) {
	router = gin.Default()

	// Enable CORS for frontend
	router.Use(cors())

	// Attach db reference for database operations
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Defaulting root url to list of commands for now
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, `
			List of endpoints:
	
			GET /npcs - Retrieve all npcs
			GET /npcs/:id - Retrieve specific npc
			POST /npcs - Create an npc
			POST /npcs/generate - Generate but do not persist NPCS
		`)
	})

	// Attach resource-specific routes
	npc.Route(router)

	err := router.Run(":" + config.PORT)
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}

	log.Printf("Server running on port %v", config.PORT)

	return
}

// Taken from https://github.com/gin-contrib/cors/issues/29#issuecomment-397859488
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
