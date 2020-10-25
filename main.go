package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/PuerkitoBio/goquery"
)

type mainTag struct{
id int
text, link string
}

func ExampleScrape() {
  var mainPage []mainTag //figure out arrays here...
  // Request the HTML page.
  res, err := http.Get("https://www.culpepperconnections.com/ss/master_index.htm")
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
  }

  // Load the HTML document
  doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
    log.Fatal(err)
  }

  // Find the review items
  doc.Find("ul#master_index").Each(func(i int, s *goquery.Selection) {
    // For each item found, get the band and title
     s.Find("li").Each(func(i int, s *goquery.Selection) {
		link := s.Find("a").Text()
		fmt.Printf("Review %d: %s \n", i, link)
	})

  })
}

func processMainLink(){
	
}

func main() {
  ExampleScrape()
}