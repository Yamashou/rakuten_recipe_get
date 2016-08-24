package main

import "github.com/PuerkitoBio/goquery"

func point(url string) []int64 {
	var points []int64
	doc, _ := goquery.NewDocument(url)
	doc.Find("dl > dd > div > p.num ").Each(func(_ int, s *goquery.Selection) {
		points = append(point, s.Text())
	})
	return points
}
