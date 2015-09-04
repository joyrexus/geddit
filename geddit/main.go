// Package main implements a basic client for the Reddit API.

package main

import (
	"fmt"
    "log"

	"github.com/joyrexus/reddit"
)

func main() {
	items, err := reddit.Get("golang")
	if err != nil {
		log.Fatal(err)
	}
		
	for _, item := range items {
		fmt.Println(item)
	}
}
