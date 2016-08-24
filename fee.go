package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func fee(url string) string {
	var i string
	doc, _ := goquery.NewDocument(url)
	doc.Find("li.icnMoney").Each(func(_ int, s *goquery.Selection) {
		i = s.Text()
	})
	return i
}

func main() {
	fmt.Print(fee("http://recipe.rakuten.co.jp/recipe/1150010609/"))
}
