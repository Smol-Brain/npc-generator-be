package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User defines the current application user (likely a DM)
type User struct {
	id    string
	name  string
	email string
}

// Route sets up endpoint routing for users
func Route(router *gin.Engine) {
	user := router.Group("/users")
	{
		user.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, getAll())
		})

		user.GET("/:id", func(c *gin.Context) {
			c.String(http.StatusOK, getOne(c.Param("id")))
		})

		user.POST("/", func(c *gin.Context) {
			c.String(http.StatusCreated, create())
		})
	}
}

func getAll() string {
	return "Getting all users"
}

func getOne(id string) string {
	return fmt.Sprintf("Getting user with id: %v", id)
}

func create() string {
	return "Creating new user"
}
