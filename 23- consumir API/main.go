package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	var clien = &http.Client{
		Timeout: 10 * time.Second,
	}
	url := "https://jsonplaceholder.typicode.com/posts"
	fmt.Println(url)

	resp, err := clien.Get(url)

	if err != nil {
		panic(err.Error())
	}

	var posts []Post

	err = json.NewDecoder(resp.Body).Decode(&posts)
	if err != nil {
		panic(err)
	}

	fmt.Println(posts)

}
