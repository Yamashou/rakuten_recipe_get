package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func time(url string) string {
	var i string
	doc, _ := goquery.NewDocument(url)
	doc.Find("time#indication_time_itemprop").Each(func(_ int, s *goquery.Selection) {
		i = s.Text()
	})
	return i
}

func main() {
	fmt.Print(time("http://recipe.rakuten.co.jp/recipe/1150010609/"))
}
