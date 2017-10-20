package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/loomnetwork/dashboard/config"
	"github.com/loomnetwork/dashboard/middleware"
	"github.com/loomnetwork/dashboard/router"
)

func Setup(db *gorm.DB, c *config.Config) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.SetDBtoContext(db))
	r.Use(middleware.SetConfigtoContext(c))
	router.Initialize(r, c)
	return r
}
