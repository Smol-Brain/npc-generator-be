package app

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"npc-generator-be/npc"
)

// InitializeServer sets up engine and routing for server requests
func InitializeServer(config Config, db *gorm.DB) (router *gin.Engine) {
	router = gin.Default()

	// Attach db reference for database operations
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Enable CORS for frontend
	router.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"https://npc-generator.netlify.com"},
			AllowMethods:     []string{"GET", "POST"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		},
	))

	// Defaulting root url to list of commands for now
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, `
			List of endpoints:
	
			GET /npcs - Retrieve all npcs
			GET /npcs/:id - Retrieve specific npc
			POST /npcs - Create an npc
			POST /npcs/generate - Generate but do not persist NPCS  (optional int req.body.npc_amount generates x npcs)
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
