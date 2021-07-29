package hardcodecast

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kelson-martins/anchor-wordpress/src/anchor"
	"github.com/kelson-martins/anchor-wordpress/src/wordpress"
)

func fetchCategories(title string) []string {
	ci := "56"
	eb := "58"
	bp := "60"

	var toReturn []string
	if strings.Contains(strings.ToLower(title), "internacional") {
		toReturn = []string{ci}
	} else if strings.Contains(strings.ToLower(title), "escovando") {
		toReturn = []string{eb}
	} else if strings.Contains(strings.ToLower(title), "byte papo") {
		toReturn = []string{bp}
	} else {
		toReturn = []string{}
	}

	return toReturn
}

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
		Format:         "audio",
		Title:          anchorData.Episodes[0].Title,
		Auth_user:      wordpressUser,
		Auth_pass:      wordpressPass,
		Content:        content,
		Categories:     fetchCategories(anchorData.Episodes[0].Title),
	}

	log.Println("[INFO] Latest Episode:" + anchorData.Episodes[0].Title)
	log.Println("[INFO] Audio link:", audioLink)
	log.Println("[INFO] Category:", post.Categories)
	log.Println("[INFO] Post into Wordpress? [Y/n]: ")

	fmt.Scan(&postConfirmation)

	if postConfirmation != "Y" {
		os.Exit(0)
	}

	return post
}
