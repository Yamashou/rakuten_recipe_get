package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func time(url string) []string {
	var i []string
	doc, _ := goquery.NewDocument(url)
	doc.Find("p > img.pc.normal").Each(func(_ int, s *goquery.Selection) {
		j, _ := s.Attr("src")
		i = append(i, j)
	})
	return i
}

func main() {
	fmt.Print(time("http://cookpad.com/cooking_basics/11206"))
}
