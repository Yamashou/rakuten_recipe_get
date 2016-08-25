package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

type process struct {
	Image     string `json:"image"`
	Operation string `json:"operation"`
}

type material struct {
	Name     []string `json:"name"`
	Quantity []string `json:"quantity"`
}

//メインのjson
type mesi struct {
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	MemberNum   string    `json:"membernum"`
	Time        string    `json:"time"`
	Fee         string    `json:"fee"`
	Explanation string    `json:"explanation"`
	Material    material  `json:"material"`
	Process     []process `json:"process"`
}
type recipe struct {
	Recipe mesi `json:"recipe"`
}

//料理名（タイトル）を返す。
func title(url string) string {
	var i string
	doc, _ := goquery.NewDocument(url)
	doc.Find("h1").Each(func(_ int, s *goquery.Selection) {
		i = s.Text()
	})
	return i
}

//何人前かを返す関数
func people(url string) string {
	var i string
	doc, _ := goquery.NewDocument(url)
	doc.Find("div > div > div > div > div > h3 > span > span").Each(func(_ int, s *goquery.Selection) {
		i = s.Text()
	})
	return i
}

//時間のをstringで返す
func time(url string) string {
	var i string
	doc, _ := goquery.NewDocument(url)
	doc.Find("time#indication_time_itemprop").Each(func(_ int, s *goquery.Selection) {
		i = s.Text()
	})
	return i
}
func fee(url string) string {
	var i string
	doc, _ := goquery.NewDocument(url)
	doc.Find("li.icnMoney").Each(func(_ int, s *goquery.Selection) {
		i = s.Text()
	})
	return i
}

//材料の量をstringの配列で返す
func materialQuantity(url string) []string {
	var quantity []string
	doc, _ := goquery.NewDocument(url)
	doc.Find("div > div > div > div > ul > li > p.amount").Each(func(_ int, s *goquery.Selection) {
		quantity = append(quantity, s.Text())
	})
	return quantity
}

//手順をstringの配列で返す
func procedure(url string) []string {
	var procedure []string
	doc, _ := goquery.NewDocument(url)
	doc.Find("li#step_box_li.stepBox > p.stepMemo").Each(func(_ int, s *goquery.Selection) {
		procedure = append(procedure, s.Text())
	})
	return procedure
}

//完成時の写真のurlをstringで返す。
func image(url string) string {
	var imgURL string
	doc, _ := goquery.NewDocument(url)
	doc.Find("div > div > span > img").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("src")
		// fmt.Println(url)
		imgURL = url
	})
	return imgURL
}

//調理時のコツについてをstringで返す。
func kotu(url string) string {
	var kotu string
	doc, _ := goquery.NewDocument(url)
	doc.Find("div.howtoPointBox.last > p").Each(func(_ int, s *goquery.Selection) {
		kotu = s.Text()
	})
	return kotu
}

//材料の名前をstringの配列で返す
func mat(url string) []string {
	var materials []string
	doc, _ := goquery.NewDocument(url)
	doc.Find("div > div > div > div > ul > li > a#material_link.name").Each(func(_ int, s *goquery.Selection) {
		materials = append(materials, s.Text())
	})
	return materials
}

//説明文を返す。
func exp(url string) string {
	var i string
	doc, _ := goquery.NewDocument(url)
	doc.Find(" div >div > div > div > p.summary").Each(func(_ int, s *goquery.Selection) {
		i = s.Text()
	})
	return i
}

//どの工程が写真を持つかを返す
func haveImages(url string) []bool {
	var test []bool
	doc, _ := goquery.NewDocument(url)

	f := doc.Find("li#step_box_li.stepBox")
	f.Each(func(_ int, s *goquery.Selection) {
		t, _ := s.Html()
		if strings.Index(t, "img") != -1 {
			test = append(test, true)
		} else {
			test = append(test, false)
		}
	})
	return test
}

//作り方内の写真を得る
func getImages(url string) []string {
	var quantity []string
	doc, _ := goquery.NewDocument(url)
	doc.Find("img#step_image.processImage").Each(func(_ int, s *goquery.Selection) {
		t, _ := s.Attr("src")
		// fmt.Println(t)
		quantity = append(quantity, t)
	})
	return quantity
}
func wait1(recipes recipe, c chan string, url string) {
	recipes.Recipe.Name = title(url)
	recipes.Recipe.Image = image(url)
	recipes.Recipe.MemberNum = people(url)
	recipes.Recipe.Explanation = exp(url)
	recipes.Recipe.Time = time(url)
	recipes.Recipe.Fee = fee(url)
	c <- "wait1 finished\n"
	// return recipes
}
func wait2(mats material, c chan string, url string) {
	mats.Quantity = materialQuantity(url)
	mats.Name = mat(url)
	c <- "wait1 finished\n"
	// return mats
}
func makejson(num string) recipe {
	// c := make(chan string)
	var recipes recipe
	var mats material
	var proc []process
	url := "http://recipe.rakuten.co.jp/recipe/"
	url = url + num
	// go wait1(recipes, c, url)
	// go wait2(mats, c, url)
	recipes.Recipe.Name = title(url)
	recipes.Recipe.Image = image(url)
	recipes.Recipe.MemberNum = people(url)
	recipes.Recipe.Explanation = exp(url)
	recipes.Recipe.Time = time(url)
	recipes.Recipe.Fee = fee(url)

	mats.Quantity = materialQuantity(url)
	mats.Name = mat(url)

	recipes.Recipe.Material.Name = mats.Name
	recipes.Recipe.Material.Quantity = mats.Quantity

	haveImage := haveImages(url)
	prImages := getImages(url)
	a := 0
	for i, operation := range procedure(url) {
		if haveImage[i] == true {
			proc = append(proc, process{
				Image:     prImages[a],
				Operation: operation,
			})
			a++
		} else {
			proc = append(proc, process{
				Image:     "",
				Operation: operation,
			})
		}
	}
	recipes.Recipe.Process = make([]process, len(proc))
	copy(recipes.Recipe.Process, proc)
	fmt.Println(recipes)
	return recipes
}

func main() {
	router := gin.Default()

	router.GET("/:num", func(c *gin.Context) {
		num := c.Param("num")
		t := makejson(num)
		c.JSON(http.StatusOK, gin.H{"recipe": t})
	})
	router.Run(":8080")
}
