package controller

import (
	"log"
	"strconv"
	"strings"

	"cheapbook/model"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/headzoo/surf.v1"
)

func Idefix(books *model.Books, s string) {
	defer wg.Done()
	bow := surf.NewBrowser()
	url := ""
	if _, err := strconv.ParseInt(s, 10, 64); err == nil {
		url = "https://www.idefix.com/search/?Q=" + s + "&ShowNotForSale=True"

	} else {
		s = strings.Replace(s, " ", "+", -1)
		url = "https://www.idefix.com/search/?Q=" + s + "&ShowNotForSale=True"

	}
	err := bow.Open(url)
	if err != nil {
		log.Println(err)
	} else {
		bow.Find(".cart-product-box-view").Each(func(index int, item *goquery.Selection) {

			title := item.Find(".box-title a").Text()
			author := item.Find(".pName a").Text()
			pub := item.Find(".manufacturerName a").Text()
			img := "aaa"
			price := item.Find("#prices").Text()
			website, _ := item.Find(".box-title a").Attr("href")

			if title != "" && price != "" {
				p := model.Book{
					Title:      title,
					Author:     author,
					Publisher:  pub,
					Img:        "http://www.idefix.com.tr" + img,
					Price:      price,
					PriceFloat: 0.0,
					WebSite:    "http://www.idefix.com.tr" + website,
					Resource:   "Idefix",
				}
				model.Add(&p, books)
			}
		})
	}
}
