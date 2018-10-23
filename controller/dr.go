package controller

import (
	"log"
	"strconv"
	"strings"


	"github.com/PuerkitoBio/goquery"
	"cheapbook/model"
	"gopkg.in/headzoo/surf.v1"
)

func Dr(books *model.Books, s string) {
	defer wg.Done()
	bow := surf.NewBrowser()
	url := ""
	if _, err := strconv.ParseInt(s, 10, 64); err == nil {
		url = "https://www.dr.com.tr/search?q="+ s

	} else {
		s = strings.Replace(s, " ", "+", -1)
		url = "https://www.dr.com.tr/search?q="+ s

	}
	err := bow.Open(url)
	if err != nil {
		log.Println(err)
	} else {
		bow.Find(".list-cell").Each(func(index int, item *goquery.Selection) {

			title := item.Find(".summary h3").Text()
			author := item.Find(".who").Text()

			pub := item.Find(".mb10").Text()

			img, _ := item.Find("figure a img").Attr("src")

			price := item.Find(".price").Text()
		//	log.Println(price)

			website, _ := item.Find(".item-name").Attr("href")
			if title != "" && price != "" {
				p := model.Book{
					Title:      title,
					Author:     author,
					Publisher:  pub,
					Img:        strings.Replace(img,"136","200",-1),
					Price:      price,
					PriceFloat: 0.0,
					WebSite:    "http://www.dr.com.tr" + website,
					Resource:   "DR",
				}
				model.Add(&p, books)
			}
		})
	}
}
