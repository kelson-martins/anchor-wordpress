package main

import (
	"fmt"

	"github.com/kelson-martins/anchor-wordpress/src/anchor"
)

const anchorURL = "https://anchor.fm/s/1157ae34/podcast/rss"

func main() {

	anchorData := anchor.ParseFeed(anchorURL)

	fmt.Println(anchorData)
}
