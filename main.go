package main

import (
	"fmt"

	"github.com/boolow5/BolScraperAPI/controllers"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	fmt.Println("Hello World")
	router := gin.Default()
	router.GET("/", controllers.Index)
	router.POST("/fetch/list", controllers.Scrape)
	router.Run()
}
