package main

import (
	fmt "fmt"
	"github.com/openfoodfacts/openfoodfacts-go"
)

func main() {
	food_api := openfoodfacts.NewClient("world", "", "")

	var barcode_number string

	for {

		fmt.Println("Insert a barcode number: ")
		fmt.Scanf("%s", &barcode_number)

		product, err := food_api.Product(barcode_number)
		if err == nil {

			fmt.Printf("%s\n", product)
		} else {
			fmt.Println("Error:", err)
		}
	}
}
