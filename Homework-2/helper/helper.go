package helper

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/FAdemoglu/homeworktwo/data"
)

func LowerCaseString(s string) string {
	return strings.ToLower(s)
}
func RangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}

func Contains(list []data.Books, searches string) {
	var searchedItems []string
	for _, v := range list {
		if strings.Contains(v.Author, searches) {
			searchedItems = append(searchedItems, v.Author)
		}
		if strings.Contains(v.BookName, searches) {
			searchedItems = append(searchedItems, v.BookName)
		}
	}
	fmt.Printf("Aranan kelimedeki yazarlar ve kitaplar %v", searchedItems)
}
