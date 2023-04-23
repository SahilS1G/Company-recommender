package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/SahilS1G/server/key"
	"github.com/SahilS1G/server/model"
)

var news []model.News

func init() {
	myUrl := fmt.Sprintf("https://newsapi.org/v2/everything?q=microsoft&apiKey=%s", key.Api_key)
	resp, err := http.Get(myUrl)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("No response from request")
	}

	var result model.News
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	news = append(news, result)
}

func GetNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// json.NewEncoder(w).Encode(news[0].Articles[0].Description)

	for i := range news[0].Articles {
		json.NewEncoder(w).Encode(news[0].Articles[i].Title)
	}
}
