package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`// omitemptyは、ゼロ値（ここではfalse）であれば、jsonを出力しない事を示す
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Gogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	}
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	data, _ = json.MarshalIndent(movies,"", "  ")
	fmt.Printf("%s\n", data)
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarchaling failed: %s", err)
	}
	fmt.Println(titles)

	// github
	Search()
}

