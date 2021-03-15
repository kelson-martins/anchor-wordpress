package main

import (
	"github.com/kelson-martins/anchor-wordpress/src/hardcodecast"
	"github.com/kelson-martins/anchor-wordpress/src/wordpress"
)

func main() {

	latestHardcodePost := hardcodecast.PostLatest("https://anchor.fm/s/1157ae34/podcast/rss", "https://hardcodecast.com/wp-json/wp/v2/posts")
	wordpress.PostArticle(latestHardcodePost)
}
