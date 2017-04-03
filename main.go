package main

import (
	"github.com/boolow5/BolScraperAPI/controllers"
	"gopkg.in/gin-gonic/gin.v1"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.GET("/", controllers.Index)
	router.POST("/fetch/list", controllers.Scrape)
	router.POST("/fetch/detail", controllers.GetDetail)
	router.Run()
}
