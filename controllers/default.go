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

func GetDetail(this *gin.Context) {
	d := models.Detail{}
	err := this.BindJSON(&d)
	if err != nil {
		this.JSON(200, gin.H{"error": err.Error()})
		return
	}
	/*
		d.Link = this.Query("link") //"http://www.jowhar.com/2017/03/17/31-qaxooti-soomaali-ah-oo-lagu-diley-duqeyn-xeebaha-dalka-yemen/"
		d.TitleSelector = this.Query("title_selector") //"#content > header > h1"
		d.DetailSelector = this.Query("detail_selector") //"#post-66495 > div"
	*/
	d.Visit()
	this.JSON(200, gin.H{"detail": d})
}
