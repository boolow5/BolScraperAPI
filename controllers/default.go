package controllers

import (
	"fmt"

	"github.com/boolow5/BolScraperAPI/models"
	"gopkg.in/gin-gonic/gin.v1"
)

func init() {
	fmt.Println("Initializing controllers")
}

func Index(this *gin.Context) {
	this.JSON(200, gin.H{"message": "Welcome to BolScraper", "more_info": "BolScraper is a simple Rest API for fetching a list of text and their links"})
}

func Scrape(this *gin.Context) {
	w := models.Website{}
	this.BindJSON(&w) // expects a json object like follows:
	//    {"name": "Jowhar", "root_url": "/home/mahdi/Downloads/webpages/jowhar.html", "selector": {"base": "article", "target_base":"h2", "target_text": "a", "target_link": "a"}}
	fmt.Println(w.Selector)
	w.Visit()

	this.JSON(200, gin.H{
		"items": w.NewsItems,
	})
}
