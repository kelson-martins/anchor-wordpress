package hardcodecast

import (
	"os"

	"github.com/kelson-martins/anchor-wordpress/src/anchor"
	"github.com/kelson-martins/anchor-wordpress/src/wordpress"
)

func PostLatest(anchorURL string, wordpressURL string) wordpress.Post {

	anchorData := anchor.ParseFeed(anchorURL)

	wordpressUser := os.Getenv("WORDPRESS_USER")
	wordpressPass := os.Getenv("WORDPRESS_PASS")

	post := wordpress.Post{
		Url:            wordpressURL,
		Status:         "draft",
		Comment_status: "open",
		Title:          anchorData.Episodes[0].Title,
		Auth_user:      wordpressUser,
		Auth_pass:      wordpressPass,
		Content:        anchorData.Episodes[0].Description,
	}

	return post
}
