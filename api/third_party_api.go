package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Response struct {
	TotalCards int    `json:"total_cards"`
	Cards      []Card `json:"data"`
}

type Card struct {
	Name string `json:"name"`
}

func ScryfallSearchCard(q string)  (*Response, error) {
	client := http.Client{}
	url := fmt.Sprintf("%scards/search?q=%s", os.Getenv("SCRYFALL_API_URL"), q)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Panicf("Failed to create request object for /GET endpoint: %v", err)
		return &Response{}, err
	}

	req.Header.Add("Content-type", "application/json; charset=utf-8")

	res, err := client.Do(req)

	if err != nil || res.StatusCode != 200 {
		log.Panicf("Failed to send HTTP request: %v", err)
		return  &Response{}, err
	}

	defer res.Body.Close()

	var responseObject Response

	err = json.NewDecoder(res.Body).Decode(&responseObject)

	if err != nil {
		log.Panicf("Failed to create request object for /GET endpoint: %v", err)
		return   &Response{}, err
	}

	return &responseObject, nil
}