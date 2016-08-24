package recipe

import "github.com/PuerkitoBio/goquery"

func mat(url string) []string {
	var materials []string
	doc, _ := goquery.NewDocument(url)
	doc.Find("div > div > div > div > ul > li > a#material_link.name").Each(func(_ int, s *goquery.Selection) {
		materials = append(materials, s.Text())
	})
	return materials
}
