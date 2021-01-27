package controller

import (
	"cheapbook/model"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/headzoo/surf.v1"
)

func Odakitap(books *model.Books, s string) {
	defer wg.Done()
	s = strings.Replace(s, " ", "+", -1)
	bow := surf.NewBrowser()
	err := bow.Open("https://www.odakitap.com/arama?q=" + s)
	if err != nil {
		log.Println(err)
	} else {
		bow.Find(".row").Each(func(index int, item *goquery.Selection) {
			title := item.Find(".plist-info h2 a").Text()
			author := item.Find("h3.author a").Text()
			pub := item.Find(".h4.store a").Text()
			img, _ := item.Find(".plist-image-wrapper img").Attr("src")
			price := item.Find(".new-price").Text()
			website, _ := item.Find(".plist-info h2 a").Attr("href")

			p := model.Book{
				Title:      title,
				Author:     author,
				Publisher:  pub,
				Img:        "https://www.odakitap.com" + img,
				Price:      price,
				PriceFloat: 0.0,
				WebSite:    "https://www.odakitap.com" + website,
				Resource:   "Odakitap",
			}
			model.Add(&p, books)
		})
	}
}
