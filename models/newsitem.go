package models

type NewsItem struct {
	ID          int    `json:"-"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	WebsiteName string `json:"website_name"`
	WebsiteURL  string `json:"website_url"`
}
