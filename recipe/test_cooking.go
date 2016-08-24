package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func procedure(url string) ([]string, []string) {
	var procedures []string
	var imgUrls []string
	doc, _ := goquery.NewDocument(url)
	img := doc.Find("li#step_box_li.stepBox > div.stepPhoto > span > img")
	doc.Find("li#step_box_li.stepBox > p.stepMemo").Each(func(_ int, s *goquery.Selection) {
		procedures = append(procedures, s.Text())
		img.Each(func(_ int, c *goquery.Selection) {
			imgurl, _ := c.Attr("src")
			imgUrls = append(imgUrls, imgurl)
		})
	})
	return procedures, imgUrls
}
