package anchor

const rss = "https://anchor.fm/s/1157ae34/podcast/rss"

type xmlStruct struct {
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

func parseFeed(url string) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("[ERROR] error fetching data from the RSS feed")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("[ERROR] error reading RSS feed")
	}

	var xmlData xmlStruct
	xml.Unmarshal(data, &xmlData)

	fmt.Println(xmlData.Episodes[0].Enclosure)
}