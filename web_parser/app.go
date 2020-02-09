package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
	Ne radi heheh
*/

// Rss : aa
type Rss struct {
	Ch []Channel `xml:"rss"`
}

// Channel : aa
type Channel struct {
	Items []Item `xml:"item"`
}

// Item : AA
type Item struct {
	Title string `xml:"title"`
}

func main() {
	resp, _ := http.Get("http://www.b92.net/info/rss/novo.xml")
	fmt.Println(resp)
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var rss Rss
	xml.Unmarshal(bytes, &rss)

	fmt.Println(rss)
}
