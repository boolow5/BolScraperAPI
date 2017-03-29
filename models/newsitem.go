package models

import (
	"errors"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type NewsItem struct {
	ID          int    `json:"-"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	WebsiteName string `json:"website_name"`
	WebsiteURL  string `json:"website_url"`
}

type Detail struct {
	ID             int    `json:"-"`
	Title          string `json:"title"`
	Link           string `json:"link"`
	Content        string `json:"content"`
	TitleSelector  string `json:"title_selector"`
	DetailSelector string `json:"detail_selector"`
}

func (this *Detail) Visit() {
	doc, err := this.fetchData()
	if err != nil {
		debug(err.Error())
	}
	if OK, err := this.parseDocument(doc); !OK || err != nil {
		if !OK {
			debug("Visiting this page didn't succeed")
		}
		if err != nil {
			debug(err.Error())
		}
	}
}

func (this *Detail) fetchData() (*goquery.Document, error) {
	debug("Fetching detail page")
	var doc *goquery.Document
	var err error
	if this.Link == "" {
		return doc, errors.New("Detail link is empty")
	}
	if strings.HasPrefix(this.Link, "http") {
		doc, err = goquery.NewDocument(this.Link)
		if err != nil {
			return doc, err
		}
	} else {
		file, err := os.Open(this.Link)
		if err != nil {
			return doc, err
		}
		doc, err = goquery.NewDocumentFromReader(file)
		if err != nil {
			return doc, err
		}
	}
	return doc, nil
}

func (this *Detail) parseDocument(doc *goquery.Document) (bool, error) {
	debug("parsing detail")
	if doc == nil {
		return false, errors.New("cannot parse and emtpy detail")
	}
	var content string = ""
	doc.Find(this.DetailSelector).Find("p").Each(func(i int, selection *goquery.Selection) {
		content += selection.Text() + "\n\n"
	})
	title := doc.Find(this.TitleSelector).Text()
	if len(title) > 1 {
		this.Title = title
		debug(title)
	}
	if len(content) > 1 {
		this.Content = content
		debug(content)
	}
	return true, nil
}
