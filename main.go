package main

import (
  "fmt"
  "gopkg.in/gin-gonic/gin.v1"
  "github.com/boolow5/BolScraperAPI/controllers"
)

func main() {
  fmt.Println("Hello World")
  router := gin.Default()
  router.GET("/", controllers.Scrape)
  router.Run()
}
