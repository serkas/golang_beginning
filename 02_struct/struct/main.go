package main

import "fmt"

//
// Structs usage:
// * representation of real-world objects (a.k.a entities)
// * organize data (value objects, DTO)
// * organize program parts (services, repositories, ...)
//

type Book struct {
	Title      string
	Author     string
	Year       int
	TotalPages int

	OrdersCount int
}

func (b Book) ShortInfo() string {
	return fmt.Sprintf("%q (%d) by %s", b.Title, b.Year, b.Author)
}

func main() {
	var book1 Book

	fmt.Printf("book1 year: %v\n", book1.Year)
	fmt.Printf("book1: %v\n", book1)

	book2 := Book{
		TotalPages: 896,
		Title:      "Dune",
		Author:     "Frank Herbert",
		Year:       1965,
	}
	//
	fmt.Printf("book2: %v\n", book2)
	fmt.Println()

	printBook(book2)
	fmt.Println(book2.ShortInfo())

	book3 := InteractiveBook{
		Book: Book{
			Title:      "Basics of Software Development for Children",
			Author:     "John Smith",
			Year:       2000,
			TotalPages: 356,
		},
		InteractivityType: BookInteractivityOnlineQuizes,
	}
	fmt.Println(book3.Book.ShortInfo())

	fmt.Println(book3.Title) // direct access to embedded fields
	fmt.Println()

	// Polimorfism
	fmt.Println("Polymorphism with interfaces:")
	printBookSummary(book2)
	printBookSummary(book3)
	fmt.Println()

	// Receivers type
	book3.SetInteractivityType(BookInteractivityPuzzle)
	fmt.Println(book3.InteractivityType)
	fmt.Println()

	// Pointers
	var v int8
	var w int8
	var p *int8
	fmt.Printf("v: %v, p: %v\n", v, p)

	v = 12
	p = &v
	fmt.Printf("v: %v, p: %v\n", v, p)

	*p += 1
	fmt.Printf("v: %v, p: %v\n", v, p)
	fmt.Println()

	p = &w
	*p += 1
	fmt.Printf("v: %v, w: %v, p: %v\n", v, w, p)

	// Structs with pointers
	shop := BookShop{
		Name:    "Milenium Books",
		Address: "61 Green street",
	}

	fmt.Println(shop.TopSellingBook) // not set yet
	//fmt.Println(shop.TopSellingBook.Title) // check it

	shop.SetTopSellingBook(&book3.Book) // limitation of embeding comparing to class ingeritance
	fmt.Println(shop.TopSellingBook)
	book3.OrdersCount += 1
	printBook(*shop.TopSellingBook)
	fmt.Println(book3.OrdersCount)

	//// Serialization -> see marshaling folder
}

func printBook(b Book) {
	fmt.Println(b.Title)
	fmt.Println(b.Author)
	fmt.Printf("Publihsed: %d\n", b.Year)
	fmt.Printf("Pages: %d\n", b.TotalPages)
	fmt.Printf("Ordes: %d\n", b.OrdersCount)
	//b.OrdersCount = 1000 // what's wrong here?
}

// interfaces are implicit (a struct implements it as soon as it has all the methods of the interface)
type BookSummary interface {
	ShortInfo() string
}

func printBookSummary(s BookSummary) {
	fmt.Println(s.ShortInfo())
}
