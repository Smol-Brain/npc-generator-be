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
		npc.DELETE("/:id", delete)
		npc.GET("/", getMany)
		npc.GET("/:id", getOne)
		npc.POST("/", create)
		npc.POST("/generate", generate)
	}
}

func getMany(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	queryParams := c.Request.URL.Query()

	var npcs []generator.Npc
	// I couldn't find a good example of Go's spread operator, will look more later
	db.Where(&generator.Npc{
		UserID: queryParams.Get("userId"),
	}).Find(&npcs)

	c.JSON(http.StatusOK, gin.H{"data": npcs, "totalCount": len(npcs)})
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

	if err := db.Create(npc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Unable to create NPC": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": npc})
}

func delete(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Delete(&generator.Npc{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Unable to delete NPC": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Deleted NPC with id: "+c.Param("id"))
}

func generate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": generateNpc()})
}
