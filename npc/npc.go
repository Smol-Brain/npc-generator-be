package npc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/lucsky/cuid"
)

// Npc defines a non-player character's traits
type Npc struct {
	ID             string
	UserID         string
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

// CreateNpcInput defines the create request body
type CreateNpcInput struct {
	UserID         string `binding:"required"`
	FirstName      string `binding:"required"`
	LastName       string `binding:"required"`
	Gender         string
	Pronouns       string
	Height         string
	Hook           string   `binding:"required"`
	Job            string   `binding:"required"`
	Languages      []string `binding:"required"`
	LifeStage      string
	NegativeTraits []string
	PositiveTraits []string
	NeutralTraits  []string
	Race           string `binding:"required"`
	Wealth         string
}

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

	var npcs []Npc
	db.Find(&npcs)

	c.JSON(http.StatusOK, gin.H{"data": npcs})
}

func getOne(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var npc Npc
	if err := db.Where("id = ?", c.Param("id")).First(&npc).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "NPC not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": npc})
}

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var requestBody CreateNpcInput
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	npc := Npc{
		ID:             "npc_" + cuid.New(), // auto-generate a cuid for now, no support for user-defined ids
		UserID:         requestBody.UserID,
		FirstName:      requestBody.FirstName,
		LastName:       requestBody.LastName,
		Gender:         requestBody.Gender,
		Pronouns:       requestBody.Pronouns,
		Height:         requestBody.Height,
		Hook:           requestBody.Hook,
		Job:            requestBody.Job,
		Languages:      requestBody.Languages,
		LifeStage:      requestBody.LifeStage,
		NegativeTraits: requestBody.NegativeTraits,
		PositiveTraits: requestBody.PositiveTraits,
		NeutralTraits:  requestBody.NeutralTraits,
		Race:           requestBody.Race,
		Wealth:         requestBody.Wealth,
	}

	db.Create(&npc)

	c.JSON(http.StatusOK, gin.H{"data": npc})
}

func generate(c *gin.Context) {
	return
}
