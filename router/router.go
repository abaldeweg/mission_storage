package router

import (
	"os"

	"github.com/abaldeweg/storage/controller/mission"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.New()
	r.SetTrustedProxies(nil)

	if os.Getenv("ENV") != "prod" {
		r.Use(gin.Logger())
	}

	r.Use(gin.Recovery())

	r.Use(headers())

	auth := r.Group("/api", checkAuth)

	// mission
	auth.GET("/mission/show", mission.Show)
	auth.POST("/mission/create", mission.Create)
	auth.PUT("/mission/update", mission.Update)
	auth.GET("/mission/export/html", mission.HtmlExport)

	r.Run(":8080")
}
