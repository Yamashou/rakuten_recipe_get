package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type kirikata struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Explanation string `json:"explanation"`
}

func getImageAndName(url string) ([]string, []string) {
	var i []string
	var j []string
	doc, _ := goquery.NewDocument(url)
	doc.Find("img.mb15 ").Each(func(_ int, s *goquery.Selection) {
		f, _ := s.Attr("alt")
		t, _ := s.Attr("src")
		i = append(i, f)
		j = append(j, t)
	})
	return i, j
}
func exp(url string) []string {
	var i []string
	doc, _ := goquery.NewDocument(url)
	doc.Find("div.txt01 ").Each(func(_ int, s *goquery.Selection) {
		i = append(i, s.Text())
	})
	return i
}
func main() {
	var kirikatas []kirikata
	name, img := getImageAndName("http://park.ajinomoto.co.jp/recipe/corner/basic/vege_cutting")
	ex := exp("http://park.ajinomoto.co.jp/recipe/corner/basic/vege_cutting")
	fmt.Print(len(ex))
	for i, s := range name {
		fmt.Print(i)
		kirikatas = append(kirikatas, kirikata{
			Name:        s,
			Image:       img[i],
			Explanation: ex[i],
		})
	}
	bytes, _ := json.Marshal(kirikatas)
	ioutil.WriteFile("./test.json", bytes, os.ModePerm)
}
