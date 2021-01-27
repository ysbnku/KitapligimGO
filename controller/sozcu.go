package controller

import (
	"log"
	"strings"

	"cheapbook/model"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/headzoo/surf.v1"
)

func Sozcu(books *model.Books, s string) {
	defer wg.Done()
	s = strings.Replace(s, " ", "+", -1)
	bow := surf.NewBrowser()
	err := bow.Open("https://www.sozcukitabevi.com/index.php?p=Products&q_field_active=0&q=" + s + "&ctg_id=&search_x=0&search_y=0&q_field=&sort_type=prs_monthly-desc&stock=1")
	if err != nil {
		log.Println(err)
	} else {
		bow.Find(".main_content").Each(func(i int, item *goquery.Selection) {
			title := "qasdasd"
			author := item.Find(".prd_view_writer span").Text()
			pub := item.Find(".publisher a").Text()
			img, _ := item.Find("#main_img img").Attr("src")
			price := item.Find("#prd_final_price_display").Text()
			website := "www.sozcukitabevi.com"

			if title != "" && price != "" {
				p := model.Book{
					Title:      title,
					Author:     author,
					Publisher:  pub,
					Img:        "https://www.sozcukitabevi.com" + img,
					Price:      price,
					PriceFloat: 0.0,
					WebSite:    website,
					Resource:   "Sözcü Kitabevi",
				}
				model.Add(&p, books)
			}

		})
	}
	// else {
	// 	item := bow.Find(".main_content")
	// 	title := item.Find(".contentHeader").Text()
	// 	author := item.Find(".prd_view_writer span").Text()
	// 	pub := item.Find(".prd_view_publisher span").Text()
	// 	img, _ := item.Find("#main_img").Attr("src")
	// 	price := item.Find("#prd_final_price_display").Text()
	// 	website := bow.Url().String()

	// 	if title != "" && price != "" {
	// 		p := model.Book{
	// 			Title:     title,
	// 			Author:    author,
	// 			Publisher: pub,
	// 			Img:       "https://www.sozcukitabevi.com" + img,
	// 			Price:     price,
	// 			WebSite:   website,
	// 			Resource:  "Sözcü Kitabevi",
	// 		}
	// 		model.Add(&p, books)
	// 	}

	// }
}
