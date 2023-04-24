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

var news []model.News

func init() {

	var query []string = []string{"microsoft"}
	// var query string = "microsoft"

	// myUrl := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&apiKey=%s", query, key.Api_key)

	filePath := "./negative_positive_keywords/negative.txt"
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	// var fileLines []string
	var result model.News

	for fileScanner.Scan() {
		// fileLines = append(fileLines, fileScanner.Text())
		query = append(query, fileScanner.Text())
		myUrl := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&apiKey=%s", strings.Join(query, "+"), key.Api_key)
		resp, err := http.Get(myUrl)
		if err != nil {
			fmt.Println("No response from request")
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("No response from request")
		}
		if err := json.Unmarshal(body, &result); err != nil {
			fmt.Println("Can not unmarshal JSON")
		}
		news = append(news, result)
		query = query[:len(query)-1]
	}

	readFile.Close()

	// for _, line := range fileLines {
	// 	fmt.Println(line)
	// }

	// fmt.Println(fileLines)

	news = append(news, result)
}

func GetNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// json.NewEncoder(w).Encode(news[0].Articles[0].Description)

	// for i := range news[0].Articles {
	json.NewEncoder(w).Encode(news[1])
	// }
}
