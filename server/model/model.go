package model

import "time"

type News struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Source      *Source   `json:"source"`
		Author      string    `json:"author"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		URL         string    `json:"url"`
		URLToImage  string    `json:"urlToImage"`
		PublishedAt time.Time `json:"publishedAt"`
		Content     string    `json:"content"`
	} `json:"articles"`
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}
type Source struct {
	ID   any    `json:"id"`
	Name string `json:"name"`
}

type NewsAPIResponse struct {
	Articles []Article `json:"articles"`
}

type SearchRequest struct {
	CompanyName string `json:"companyName"`
}
