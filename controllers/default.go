package controllers

import (
  "fmt"
  "gopkg.in/gin-gonic/gin.v1"
  "github.com/boolow5/BolScraperAPI/models"
)

func init() {
  fmt.Println("Initializing controllers")
}

func Scrape(this *gin.Context) {
  w := models.Website{}
  w.Name = "Jowhar"
  w.RootURL = "/home/mahdi/Downloads/webpages/jowhar.html"
  w.Selector = &models.Selector{
    Base: "#content > div.box-news > section.left-column > article",
    TargetBase: "h2",
    TargetText: "a",
    TargetLink: "a",
  }
  w.Visit()
  this.JSON(200, gin.H{
    "website": w,
  })
}
