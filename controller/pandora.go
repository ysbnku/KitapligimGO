package controller

import (
	"log"
	"strconv"
	"strings"

	"cheapbook/model"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/headzoo/surf.v1"
)

func Pandora(books *model.Books, s string) {
	defer wg.Done()
	bow := surf.NewBrowser()
	url := ""
	if _, err := strconv.ParseInt(s, 10, 64); err == nil {
		url = "https://www.pandora.com.tr/Arama/?type=9&kitapadi=&yazaradi=&yayinevi=&isbn=" + s + "&dil=&siteid=&kategori=&sirala=0"

	} else {
		s = strings.Replace(s, " ", "+", -1)
		url = "https://www.pandora.com.tr/Arama/?type=9&kitapadi=&yazaradi=&yayinevi=&isbn=" + s + "&dil=&siteid=&kategori=&sirala=0"
	}
	err := bow.Open(url)
	if err != nil {
		log.Println(err)
	} else {
		bow.Find(".indirimVar").Each(func(index int, item *goquery.Selection) {

			title := item.Find(".edebiyatIsim strong").Text()
			author := item.Find(".edebiyatYazar a").Text()

			pub := item.Find(".edebiyatYayinEvi").Text()

			img, _ := item.Find(".coverWrapper img").Attr("src")
			price := item.Find(".indirimliFiyat").Text()

			website, _ := item.Find(".edebiyatIsim a").Attr("href")
			if title != "" && price != "" {
				p := model.Book{
					Title:      title,
					Author:     author,
					Publisher:  pub,
					Img:        "https://www.pandora.com.tr" + img,
					Price:      substring(price),
					PriceFloat: 0.0,
					WebSite:    "https://www.pandora.com.tr" + website,
					Resource:   "Pandora",
				}
				model.Add(&p, books)
			}
		})
	}
}

func substring(c string) string {

	runes := []rune(c)
	price := string(runes[13:len(runes)])
	return price

}
