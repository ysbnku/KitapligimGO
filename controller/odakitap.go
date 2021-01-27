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
		log.Println("aaaaaaaa")
	} else {
		bow.Find(".row").Each(func(index int, item *goquery.Selection) {
			title := "kitap"
			author := "kitapcııı"
			pub := "kitaplaaaar"
			img, _ := item.Find(".plist-image-wrapper img").Attr("src")
			price := "aaaaaaaa"
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
