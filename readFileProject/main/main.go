package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//reading Golang_Test.txt file.
	fileContent, err := ioutil.ReadFile("C:/Users/Krrish/go/FindTop10Words/readFileProject/main/Golang_Test.txt")
	if err != nil {
		log.Fatalf("An Error Occured in file reading %v", err)
	}
	//fmt.Println(string(fileContent))

	// preparing post request data
	postRequestBodyData, _ := json.Marshal(map[string]string{
		"content": string(fileContent),
	})
	requestBodyBuffer := bytes.NewBuffer(postRequestBodyData)

	// Making HTTP request to our service
	resp, err := http.Post("http://localhost:8080/top10MostUsedWords", "application/json", requestBodyBuffer)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	//Read the response body
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	responseDataJson := string(responseData)
	log.Println(responseDataJson)

}
