package npc

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Npc defines a non-player character's traits
type Npc struct {
	ID             string
	FirstName      string
	LastName       string
	Gender         string
	Pronouns       string
	Height         string
	Hook           string
	Job            string
	Languages      []string `gorm:"type:varchar(64)[]"`
	LifeStage      string
	NegativeTraits []string `gorm:"type:varchar(64)[]"`
	PositiveTraits []string `gorm:"type:varchar(64)[]"`
	NeutralTraits  []string `gorm:"type:varchar(64)[]"`
	Race           string
	Wealth         string
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

func delete(id string) string {
	return fmt.Sprintf("Deleting NPC with id: %v", id)
}
