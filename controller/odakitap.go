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
		bow.Find(".plist-item").Each(func(i int, item *goquery.Selection) {
			title := "Deneme"
			author := item.Find(".l-owner h3").Text()
			pub := item.Find(".l-owner h4").Text()
			img, _ := item.Find(".plist-image-wrapper img").Attr("src")
			price := "22"
			website, _ := item.Find(".plist-image-wrapper a").Attr("href")

			if title != "" && price != "" {
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
			}
		})
	}
}
