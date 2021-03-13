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
}

func fetchCategory(title string) int {
	ci := 56
	eb := 58

	var toReturn int

	if strings.Contains(title, "Internacional") {
		toReturn = ci
	} else if strings.Contains(title, "Escovando") {
		toReturn = eb
	} else {
		toReturn = 0
	}

	return toReturn
}

func PostArticle(post Post) {

	categoryString := ""

	category := fetchCategory(post.Title)
	postUrl := post.Url

	content := strings.Replace(post.Content, "<p>", "", -1)
	content = strings.Replace(content, "</p>", "", -1)
	content = strings.Replace(content, "<strong>", "", -1)
	content = strings.Replace(content, "</strong>", "", -1)
	content = strings.Replace(content, "\n", "", -1)

	if category > 0 {
		categoryString = fmt.Sprintf(`"%v"`, category)
	}

	dataString := fmt.Sprintf(`
		{
			"status":         "%v",
			"comment_status": "%v",
			"title":          "%v",
			"content":        "%v",
			"categories": [ %v ]
		}	
	`, post.Status, post.Comment_status, post.Title, content, categoryString)

	data := []byte(dataString)

	fmt.Println(dataString)

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
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	fmt.Printf("%s\n", body)
}
