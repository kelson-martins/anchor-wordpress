package hardcodecast

import (
	"fmt"
	"os"

	"github.com/kelson-martins/anchor-wordpress/src/anchor"
	"github.com/kelson-martins/anchor-wordpress/src/wordpress"
)

func PostLatest(anchorURL string, wordpressURL string) wordpress.Post {

	var postConfirmation string

	anchorData := anchor.ParseFeed(anchorURL)

	wordpressUser := os.Getenv("WORDPRESS_USER")
	wordpressPass := os.Getenv("WORDPRESS_PASS")

	// appending the audio link
	audioLink := "Link: " + anchorData.Episodes[0].Enclosure.URL
	content := fmt.Sprint(anchorData.Episodes[0].Description + " " + audioLink)

	post := wordpress.Post{
		Url:            wordpressURL,
		Status:         "draft",
		Comment_status: "open",
		Title:          anchorData.Episodes[0].Title,
		Auth_user:      wordpressUser,
		Auth_pass:      wordpressPass,
		Content:        content,
	}

	fmt.Println("Latest Episode: " + anchorData.Episodes[0].Title)
	fmt.Printf("Post into Wordpress? [Y/n]: ")

	fmt.Scan(&postConfirmation)

	if postConfirmation != "Y" {
		os.Exit(0)
	}

	return post
}
