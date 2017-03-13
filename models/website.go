package models

import (
  "os"
  "strings"
  "errors"
  "github.com/PuerkitoBio/goquery"
)

type Website struct {
  ID int
  Name string
  Title string
  RootURL string
  Selector *Selector
  NewsItems []*NewsItem
}

type Selector struct {
  Base string
  TargetBase string
  TargetText string
  TargetLink string
}

func (site *Website) Visit() {
  doc, err := site.FetchData()
  if err != nil {
    debug(err.Error())
  }
  site.parseDocument(doc)
  for item := range site.NewsItems {
    debug("item: %s\nLink:\t%s\n", site.NewsItems[item].Title, site.NewsItems[item].Link)
  }
}

func (site Website) FetchData() (*goquery.Document, error) {
  debug("fetching data...")
  var doc *goquery.Document
  var err error
  // if emtpy
  if site.RootURL == "" {
    return doc, errors.New("Root URL is emtpy")
  }
  // check the site url if its http or file name
  if strings.HasPrefix(site.RootURL, "http") {
    // it's a url for website
    debug("this url is for website")
    debug("calling website server")
    doc, err = goquery.NewDocument(site.RootURL)
    if err != nil {
      return doc, err
    }
  } else {
    // it's file
    debug("this url is for html file")
    file, err := os.Open(site.RootURL)
    if err != nil {
      return doc, err
    }
    debug("opening file")
    doc, err = goquery.NewDocumentFromReader(file)
    if err != nil {
      return doc, err
    }
  }
  debug("fetched document")
  return doc, nil
}

func (site *Website) parseDocument(doc *goquery.Document) (bool, error) {
  debug("Parsing document")
  // check doc is nil
  if doc == nil {
    return false, errors.New("cannot parse an emtpy document")
  }
  // parse news items from the document
  var added, skipped int
  doc.Find(site.Selector.Base).Each(func(i int, s *goquery.Selection) {
    section := s.Find(site.Selector.TargetBase)
    //debug(section.Text())
    link, _ := section.Find(site.Selector.TargetLink).Attr("href")
    debug(link)
    text := section.Find(site.Selector.TargetText).Text()
    debug(text)
    if len(link) > 0 && len(text) > 0 {
      site.NewsItems = append(site.NewsItems, &NewsItem{Title: text, Link: link, WebsiteName: site.Name, WebsiteURL: site.RootURL})
      added += 1
    } else {
      skipped += 1
    }
  })
  debug("added: %d\tskipped: %d\n", added, skipped)
  return len(site.NewsItems) != 0, nil
}
