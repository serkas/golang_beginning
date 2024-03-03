package main

type BookShop struct {
	Name           string
	Address        string
	TopSellingBook *Book
}

func (s *BookShop) SetTopSellingBook(b *Book) {
	s.TopSellingBook = b
}
