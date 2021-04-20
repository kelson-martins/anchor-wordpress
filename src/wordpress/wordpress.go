package wordpress

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http"
	"strings"
	"time"
)

type Post struct {
	Url            string
	Status         string
	Comment_status string
	Title          string
	Auth_user      string
	Auth_pass      string
	Categories     []string
	Content        string
	Slug           string
	Format         string
}

func getCategories(categories []string) string {

	var parsedCategories string

	for i, v := range categories {

		if i > 0 {
			parsedCategories = parsedCategories + ","
		}

		parsedCategories = fmt.Sprint(parsedCategories, '"', v, '"')
	}

	return parsedCategories

}

func PostArticle(post Post) {

	postUrl := post.Url

	content := strings.Replace(post.Content, "<p>", "", -1)
	content = strings.Replace(content, "</p>", "", -1)
	content = strings.Replace(content, "<strong>", "", -1)
	content = strings.Replace(content, "</strong>", "", -1)
	content = strings.Replace(content, "\n", "", -1)

	dataString := fmt.Sprintf(`
		{
			"status":         "%v",
			"comment_status": "%v",
			"title":          "%v",
			"content":        "%v",
			"format":         "%v",
			"categories": [ %v ]
		}	
	`, post.Status, post.Comment_status, post.Title, content, post.Format, getCategories(post.Categories))

	data := []byte(dataString)

	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(data))

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(post.Auth_user, post.Auth_pass)

	// set client timeout
	client := &http.Client{Timeout: time.Second * 10}

	// send request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	} else {
		fmt.Println("Article posted")
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

}
