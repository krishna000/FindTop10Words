package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

func main() {
	fmt.Println("trie service starting")

	handleRequests()

}

func serverTop10MostUsedWords(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	reqData := &RequestData{}
	err := json.Unmarshal(reqBody, reqData)
	if err != nil {
		log.Fatal(err)
	}

	resData := ResponseData{Top10Words: convAns(getTop10MostUsedWordsFromContent(reqData.Content))}
	resDataJsonString, err := json.Marshal(resData)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintf(w, string(resDataJsonString))
	if err != nil {
		log.Fatal(err)
	}
}

func convAns(contents *MinHeap) []WordCount {
	var ans = make([]WordCount, 0)

	for _, content := range *contents {
		ans = append(ans, WordCount{content.Word, content.Count})
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i].Count > ans[j].Count
	})
	return ans
}

func handleRequests() {
	http.HandleFunc("/top10MostUsedWords", serverTop10MostUsedWords)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type RequestData struct {
	Content string
}

type ResponseData struct {
	Top10Words []WordCount
}
