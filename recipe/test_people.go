package recipe

import "github.com/PuerkitoBio/goquery"

func people(url string) string {
	var i string
	doc, _ := goquery.NewDocument(url)
	doc.Find("div > div > div > div > div > h3 > span > span").Each(func(_ int, s *goquery.Selection) {
		i = s.Text()
	})
	return i
}
