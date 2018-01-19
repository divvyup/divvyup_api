package examples

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// InsertReceiptData is
func InsertReceiptData() {
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
		fmt.Println(len(item))
		if len(item) < 8 && !IsLetter(item) {
			foo := reg.ReplaceAllString(item, "")
			price, _ := strconv.ParseFloat(foo, 10)
			prices = append(prices, price)
		} else if !strings.Contains(item, "@") {
			items = append(items, item)
		}
	}
	fmt.Println(len(prices))
	fmt.Println(len(items))
	for index, item := range items {
		fmt.Print(item + ": $")
		fmt.Println(prices[index])
	}
	total := 0.0
	for _, price := range prices {
		total += price
	}
	// because we add the subtotal divide in half
	total = total / 2
	fmt.Println("-----------------------------------")
	fmt.Print("Calculated total: $")
	fmt.Println(total)

	os.Exit(1)
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
