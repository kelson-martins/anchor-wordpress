package anchor

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

type anchorStruct struct {
	Title       string    `xml:"channel>title"`
	Episodes    []episode `xml:"channel>item"`
	Description string    `xml:"channel>description"`
}

type episode struct {
	Title       string    `xml:"title"`
	Description string    `xml:"description"`
	Enclosure   enclosure `xml:"enclosure"`
}

type enclosure struct {
	URL string `xml:"url,attr"`
}

func ParseFeed(url string) anchorStruct {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("[ERROR] error fetching data from the RSS feed")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("[ERROR] error reading RSS feed")
	}

	var anchorData anchorStruct
	xml.Unmarshal(data, &anchorData)

	return anchorData
}
