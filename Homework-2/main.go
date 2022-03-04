package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/FAdemoglu/homeworktwo/data"
	"github.com/FAdemoglu/homeworktwo/helper"
	valid "github.com/asaskevich/govalidator"
)

var books []data.Books

func init() {
	minPrice := 10
	maxPrice := 50
	ISBNNumber := helper.RangeIn(10000000, 99999999)
	book1, err := data.NewBook(1,
		"Sapiens",
		ISBNNumber,
		ISBNNumber,
		250,
		rand.Intn(maxPrice-minPrice)+minPrice,
		rand.Intn(maxPrice-minPrice)+minPrice,
		"Yual Noah Harrari", false)
	if err != nil {
		fmt.Println(err.Error())
	}

	books = append(books, *book1)
}

func main() {
	args := os.Args
	if len(args) == 1 {
		projectName := path.Base(args[0])
		fmt.Printf("%s uygulamasında kullanabileceğiniz komutlar : \n search=> arama işlemi için \n list=> listeleme işlemi için \n delete=> silme işlemi için \n buy=> kitabı satın almak için\n", projectName)
		return
	}

	argument := helper.LowerCaseString(args[1])

	if argument == "search" && len(args) < 3 {
		fmt.Printf("Aramak istediğiniz kitabı veya yazarı girmelisiniz!")
	} else if (argument == "buy" || argument == "delete") && len(args) >= 2 {
		for i := 2; i < len(args); i++ {
			if !valid.IsInt(args[i]) {
				fmt.Printf("Numerik bir değer girmelisiniz\n")
				return
			}
		}
		switch argument {
		case "search":
			Search(books, strings.Join(args[2:], " "))
		case "list":
			List(books)
		case "buy":
			if len(args) < 4 {
				fmt.Printf("Diğer komutları da girmelisiniz")
			} else {
				firstArgument, _ := strconv.Atoi(args[2])
				secondArgument, _ := strconv.Atoi(args[3])
				Buy(books, firstArgument, secondArgument)
			}
		case "delete":
			if len(args) < 3 {
				fmt.Printf("Diğer komutları da girmelisiniz")
			} else {
				firstArgument, _ := strconv.Atoi(args[2])
				Delete(books, firstArgument)
			}
		default:
			fmt.Printf("Yanlış bir komut girdiniz.")
		}
	}
}

func Search(books []data.Books, search string) {
	helper.Contains(books, search)
}

func List(books []data.Books) {
	var listBooks []data.Books
	for _, v := range books {
		if v.IsDeleted == false {
			listBooks = append(listBooks, v)
		}
	}
	fmt.Printf("%v", listBooks)
}

func Buy(books []data.Books, Id int, Count int) {
	for _, v := range books {
		if Id == v.Id && v.IsDeleted == false {
			v.Sell(v.Id, Count)
		} else {
			fmt.Printf("Bu ürün kaldırılmış veya bu Id'ye sahip bir ürün bulunmamaktadır.\n")
		}
	}
	fmt.Println(books)
}

func Delete(books []data.Books, Id int) {
	for _, v := range books {
		if v.IsDeleted == true {
			fmt.Printf("Bu kitap silinmiş.")
		} else {
			if v.Id == Id {
				v.Delete(v.Id)
			} else {
				fmt.Printf("Bu Id ile bir kitap bulunmamaktadır.")
			}
		}
	}
}
