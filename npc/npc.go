package npc

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Npc defines a non-player character's traits
// I would prefer snake_case, but I guess Go has other opinions :3
type Npc struct {
	firstName, lastName                           string
	gender, pronouns                              string
	height                                        string
	hook                                          string
	job                                           string
	languages                                     []string
	lifeStage                                     string
	negativeTraits, neutralTraits, positiveTraits []string
	race                                          string
	wealth                                        string
}

// Route sets up endpoint routing for NPCs
func Route(router *gin.Engine) {
	npc := router.Group("/npcs")
	{
		npc.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, getAll())
		})

		npc.GET("/:id", func(c *gin.Context) {
			c.String(http.StatusOK, getOne(c.Param("id")))
		})

		npc.POST("/", func(c *gin.Context) {
			c.String(http.StatusCreated, create())
		})
	}
}

func getAll() string {
	return "Getting all NPCs"
}

func getOne(id string) string {
	return fmt.Sprintf("Getting NPC with id: %v", id)
}

func create() string {
	return "Creating new NPC"
}
