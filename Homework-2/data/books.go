package data

import "fmt"

type Books struct {
	Id         int
	BookName   string
	StockCode  int
	ISBNNumber int
	PageNumber int
	Price      int
	StockCount int
	Author     string
	IsDeleted  bool
}

type Deletable interface {
	Delete(id int)
}

func (b *Books) Delete(id int) {
	if b.Id == id {
		b.IsDeleted = true
		fmt.Printf("Kitap başarılı bir şekilde silindi.")
	}
}

func NewBook(Id int, BookName string, StockCode int, ISBNNumber int, PageNumber int, Price int, StockCount int, Author string, IsDeleted bool) (*Books, error) {
	book := &Books{
		Id:         Id,
		BookName:   BookName,
		StockCode:  StockCode,
		ISBNNumber: ISBNNumber,
		PageNumber: PageNumber,
		Price:      Price,
		StockCount: StockCount,
		Author:     Author,
		IsDeleted:  IsDeleted,
	}
	return book, nil
}

func (b *Books) Sell(id int, i int) {
	b.StockCount -= i
	fmt.Printf("Kitap başarılı bir şekilde satıldı")
	fmt.Println(b)
}
