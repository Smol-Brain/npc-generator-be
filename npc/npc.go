package npc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucsky/cuid"
	"gorm.io/gorm"

	generator "github/Smol-Brain/npc-generator-be"
)

// Route sets up endpoint routing for NPCs
func Route(router *gin.Engine) {
	npc := router.Group("/npcs")
	{
		npc.GET("/", getMany)
		npc.GET("/:id", getOne)
		npc.POST("/", create)
		npc.POST("/generate", generate)
	}
}

func getMany(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var npcs []generator.Npc
	db.Find(&npcs)

	c.JSON(http.StatusOK, gin.H{"data": npcs})
}

func getOne(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var npc generator.Npc
	if err := db.Where("id = ?", c.Param("id")).First(&npc).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "NPC not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": npc})
}

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var npc generator.Npc
	if err := c.ShouldBindJSON(&npc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	// Auto-assign id regardless if present in request body
	npc.ID = "npc_" + cuid.New()

	db.Create(&npc)

	c.JSON(http.StatusOK, gin.H{"data": npc})
}

func generate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": generateNpc()})
}
