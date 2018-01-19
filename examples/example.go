package examples

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/domtheporcupine/divvyup_api/models"
)

// ExampleItems will be used in our demo version
var ExampleItems = []models.Item{}

// ParseReceiptData is
func ParseReceiptData() {
	data, err := ioutil.ReadFile("examples/googlevision.txt")

	if err != nil {
		log.Fatal(err)
	}

	var prices []float64
	var items []string

	reg, err := regexp.Compile("[^0-9.]+")
	if err != nil {
		log.Fatal(err)
	}

	iterms := strings.Split(string(data), "\n")

	for _, item := range iterms {
		if len(item) < 8 && !IsLetter(item) {
			foo := reg.ReplaceAllString(item, "")
			price, _ := strconv.ParseFloat(foo, 10)
			prices = append(prices, price)
		} else if !strings.Contains(item, "@") {
			items = append(items, item)
		}
	}

	for index, item := range items {
		nItem := models.Item{}
		nItem.Price = prices[index]
		nItem.Name = item
		ExampleItems = append(ExampleItems, nItem)
	}

	// os.Exit(1)
}

/*
	IsLetter is a function to tell us if
	the string s contains only letters
*/
func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

/*
	GetExampleItems returns a list of Items
	parsed from the google vision response to
	the picture of the receipt in this directory
*/
func GetExampleItems() []models.Item {
	return ExampleItems
}
