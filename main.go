package main

import (
	fmt "fmt"

	"github.com/openfoodfacts/openfoodfacts-go"
)

func retrieve_allergens(product *openfoodfacts.Product) {

	fmt.Printf("%s\n", product.Allergens)

	if len(product.Allergens) <= 0 {
		fmt.Println("Null allergens")
	}

}

func main() {
	food_api := openfoodfacts.NewClient("world", "", "")

	var barcode_number string

	for {

		fmt.Println("Insert a barcode number: ")
		fmt.Scanf("%s", &barcode_number)

		product, err := food_api.Product(barcode_number)
		if err == nil {
			retrieve_ingredients(product)
		} else {
			fmt.Println("Error:", err)
		}
	}
}
