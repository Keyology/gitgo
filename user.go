package main

import (
	"encoding/json"
	"fmt"

	// "time"
	// "fmt"
	// "io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	apiURL = "https://api.microlink.io?url="
)


func GetLink() {
	newData := "https://www.youtube.com/watch?v=x4Xt3P7FQ2M&t=924s"
	var payload Schema

	resp, err := http.Get(apiURL + newData)

	if err != nil {

		log.Fatalf("Error retrieving data: %s\n", err)

	}

	json.NewDecoder(resp.Body).Decode(&payload)
	json.NewEncoder(os.Stdout).Encode(payload)
	fmt.Println("This is before close", payload)

	resp.Body.Close()
	fmt.Println("This is after close", payload)

}
func main(){
	GetLink()
}
// func getUsers(name string) User {
// 	resp, err := http.Get(apiURL + userEndpoint + name)
// 	if err != nil {
// 		log.Fatalf("Error retrieving data: %s\n", err)
// 	}

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalf("Error reading data: %s\n", err)
// 	}

// 	var user User
// 	json.Unmarshal(body, &user)
// 	return user
// }

// Microlink API follows JSend schema format
// See more: https://labs.omniti.com/labs/jsend

type Schema struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    *Data  `json:"data,omitempty"`
}

type Data struct {
	Author      string    `json:"author,omitempty"`
	Lang        string    `json:"lang,omitempty"`
	Date        string    `json:"date,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Publisher   string    `json:"publisher,omitempty"`
	URL         string    `json:"url"`
	Image       *ImageURL `json:"image,omitempty"`
	Video       *ImageURL `json:"video,omitempty"`
	Logo        *ImageURL `json:"logo,omitempty"`
	Screenshot  *ImageURL `json:"screenshot,omitempty"`
}

type ImageURL struct {
	Width            int      `json:"width,omitempty"`
	Height           int      `json:"height,omitempty"`
	Size             int      `json:"size,omitempty"`
	SizePretty       string   `json:"size_pretty,omitempty"`
	Duration         int      `json:"duration,omitempty"`
	DurationPretty   string   `json:"duration_pretty,omitempty"`
	Type             string   `json:"type,omitempty"`
	URL              string   `json:"url"`
	Palette          []string `json:"palette,omitempty"`
	BackgroundColor  string   `json:"background_color,omitempty"`
	AlternativeColor string   `json:"alternative_color,omitempty"`
	Color            string   `json:"color,omitempty"`
}
