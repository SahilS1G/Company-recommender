package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"

	"io"
	"net/http"
	"os"

	"strings"

	"github.com/SahilS1G/server/key"
	"github.com/SahilS1G/server/model"
)

var newsResponse model.NewsAPIResponse

func init() {
	
	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=Infosys&apiKey=%s", key.Api_key)
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
	fmt.Println(newsResponse.Articles[1].Description)
	// Loop over all articles and categorize them as positive or negative
	var positiveArticles []model.Article
	var negativeArticles []model.Article
	for _, article := range newsResponse.Articles {
		sentiment := getNewsSentiment(article.Title + " " + article.Description)
		if sentiment == "Positive" {
			positiveArticles = append(positiveArticles, article)
		} else if sentiment == "Negative" {
			negativeArticles = append(negativeArticles, article)
		}
	}

	// Print the results
	fmt.Println("Positive articles:")
	for _, article := range positiveArticles {
		fmt.Println(article.Title)
	}

	fmt.Println("Negative articles:")
	for _, article := range negativeArticles {
		fmt.Println(article.Title)
	}

}

func getNewsSentiment(article string) string {
	var positiveWords []string
	var negativeWords []string

	filePath := "./negative_positive_keywords/positive.txt"
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		positiveWords = append(positiveWords, strings.ToLower(fileScanner.Text()))

	}

	readFile.Close()
	filePath2 := "./negative_positive_keywords/negative.txt"
	readFile2, err := os.Open(filePath2)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner2 := bufio.NewScanner(readFile2)
	fileScanner2.Split(bufio.ScanLines)

	for fileScanner2.Scan() {
		negativeWords = append(negativeWords, strings.ToLower(fileScanner2.Text()))

	}

	readFile.Close()

	articleWords := strings.Split(strings.ToLower(article), " ")
	positiveCount := 0
	negativeCount := 0

	for _, word := range articleWords {
		if contains(positiveWords, word) {
			positiveCount++
		} else if contains(negativeWords, word) {
			negativeCount++
		}
	}

	if positiveCount > negativeCount {
		return "Positive"
	} else if negativeCount > positiveCount {
		return "Negative"
	} else {
		return "Neutral"
	}
}

func contains(words []string, word string) bool {
	for _, w := range words {
		if w == word {
			return true
		}
	}
	return false
}

func GetNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// json.NewEncoder(w).Encode(news[0].Articles[0].Description)

	// for i := range news[0].Articles {
	json.NewEncoder(w).Encode(newsResponse)

	// }
}
