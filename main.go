package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kelson-martins/anchor-wordpress/anchor"
)



func main() {
	anchor.parseFeed(rss)
}
