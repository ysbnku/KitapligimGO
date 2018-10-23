package controller

import (
	"log"
	"strconv"
	"strings"


	"github.com/PuerkitoBio/goquery"
	"cheapbook/model"
	"gopkg.in/headzoo/surf.v1"
)

func KitapYurdu(books *model.Books, s string) {
	defer wg.Done()
	bow := surf.NewBrowser()
	url := ""
	if _, err := strconv.ParseInt(s, 10, 64); err == nil {
		url = "https://www.kitapyurdu.com/index.php?route=product/search&filter_name="+ s + "&limit=100"



	} else {
		s = strings.Replace(s, " ", "+", -1)
		url = "https://www.kitapyurdu.com/index.php?route=product/search&filter_name="+ s + "&limit=100"

	}
	err := bow.Open(url)
	if err != nil {
		log.Println(err)
	} else {
		bow.Find(".grid_7 omega").Each(func(index int, item *goquery.Selection) {

			title := item.Find(".name span").Text()
			author := item.Find(".author span").Text()

			pub := item.Find(".publisher span").Text()

			img, _ := item.Find(".image img").Attr("src")

			price := item.Find(".price-new span").Text()

			website, _ := item.Find(".cover a").Attr("href")
			if title != "" && price != "" {
				p := model.Book{
					Title:      title,
					Author:     author,
					Publisher:  pub,
					Img:        img,
					Price:      price,
					PriceFloat: 0.0,
					WebSite:     website,
					Resource:   "Kitap Yurdu",
				}
				model.Add(&p, books)
			}
		})
	}
}
