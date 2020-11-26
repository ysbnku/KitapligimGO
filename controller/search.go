package controller

import (
	"sync"

	"cheapbook/model"
)

var wg sync.WaitGroup

func Search(books *model.Books, s string) *model.Books {
	wg.Add(7)
	go Idefix(books, s)
	// go Odakitap(books, s)
	go Pandora(books, s)
	go Hepsiburada(books, s)
	go Sozcu(books, s)
	go Dr(books, s)
	// go KitapYurdu(books, s)
	wg.Wait()
	return books
}
