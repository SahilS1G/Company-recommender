package controllers

import (
	"encoding/json"
	"fmt"

	"io"
	"net/http"

	"github.com/SahilS1G/server/key"
	"github.com/SahilS1G/server/model"
	"github.com/jonreiter/govader"
)

var CompanyName string

var newsResponse model.NewsAPIResponse
var positiveArticles []model.Article
var negativeArticles []model.Article

func negative_positive(company string) {
	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&apiKey=%s", company, key.Api_key)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request to News API:", err)
		return
	}
	defer resp.Body.Close()

	// Parse the response JSON into a NewsAPIResponse struct
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response from News API:", err)
		return
	}
	err = json.Unmarshal(body, &newsResponse)
	if err != nil {
		fmt.Println("Error parsing response from News API:", err)
		return
	}

	// Loop over all articles and categorize them as positive or negative

	for _, article := range newsResponse.Articles {
		sentiment := getNewsSentiment(article.Title + " " + article.Description)
		if sentiment == "Positive" {
			positiveArticles = append(positiveArticles, article)

		} else if sentiment == "Negative" {

			negativeArticles = append(negativeArticles, article)
		}
	}

}
func GetNews(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r)
	w.Header().Set("Content-Type", "application/json")
	company := "tesla"
	negative_positive(company)
	json.NewEncoder(w).Encode(
		model.NewsAPIResponse{
			Articles: newsResponse.Articles,
		},
	)
}

func getNewsSentiment(article string) string {

	vader := govader.NewSentimentIntensityAnalyzer()
	scores := vader.PolarityScores(article)
	if scores.Compound > 0 {
		return "Positive"
	} else if scores.Compound < 0 {
		return "Negative"
	} else {
		return "Neutral"
	}

}

func enableCors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
}

func GetPositveNews(w http.ResponseWriter, r *http.Request) {

	enableCors(w, r)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		positiveArticles,
	)

}

func GetNegativeNews(w http.ResponseWriter, r *http.Request) {

	enableCors(w, r)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(
		negativeArticles,
	)
}

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	enableCors(w, r)

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body into a SearchRequest struct
	var req model.SearchRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Access the company name from the request
	companyName := req.CompanyName

	// Do something with the company name (e.g., store it in a variable, perform processing, etc.)
	fmt.Println("Received search request for company:", companyName)

	// Send a response back to the client
	response := map[string]string{
		"message": "Search request received",
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	negative_positive(companyName)
}
