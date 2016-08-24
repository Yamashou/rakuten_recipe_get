package recipe

import "github.com/PuerkitoBio/goquery"

func kotu(url string) string {
	var kotu string
	doc, _ := goquery.NewDocument(url)
	doc.Find("div.howtoPointBox.last > p").Each(func(_ int, s *goquery.Selection) {
		kotu = s.Text()
	})
	return kotu
}
