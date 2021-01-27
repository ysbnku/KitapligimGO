package controller

import (
	"log"
	"strconv"
	"strings"

	"cheapbook/model"

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
	} else if _, ok := strconv.ParseFloat(s, 64); ok != nil {
		bow.Find(".plist-item").Each(func(i int, item *goquery.Selection) {
			title := item.Find(".plist-info h2").Text()
			author := item.Find(".l-owner h3").Text()
			pub := item.Find(".l-owner h4").Text()
			img, _ := item.Find(".plist-image-wrapper img").Attr("src")
			price := item.Find(".new-price").Text()
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
	} else {
		item := bow.Find(".main-content")

		title := item.Find(".pd-name").Text()
		author := item.Find(".pd-owner a").Text()
		pub := item.Find(".pd-publisher a span").Text()
		img, _ := item.Find("#main_img").Attr("src")
		price := item.Find("#prd_final_price_display").Text()
		website := bow.Url().String()

		if title != "" && price != "" {
			p := model.Book{
				Title:     title,
				Author:    author,
				Publisher: pub,
				Img:       "https://www.odakitap.com" + img,
				Price:     price,
				WebSite:   website,
				Resource:  "Odakitap",
			}
			model.Add(&p, books)
		}

	}
}
