package global

import (
	"gorm.io/gorm"

	"npc-generator-be/config"
)

// Conf is a global var that holds configuration values
var Conf config.Config

// DB is the database instance
var DB *gorm.DB
