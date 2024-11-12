package main

import (
	"fmt"
	"time"

	"github.com/anthonyoliai/trigger-products-go/storage"
)

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/product-store?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := storage.New(dsn)
	if err != nil {
		panic(err)
	}
	if err := db.CreateTriggers(); err != nil {
		panic(err)
	}

	releaseDate := time.Date(2024, 10, 8, 0, 0, 0, 0, time.Local)
	phone := storage.Product{Name: "Galaxy S20", Country: "Netherlands", Price: 1000.0, ReleaseDate: releaseDate}
	laptop := storage.Product{Name: "Macbook Air M1 2019", Country: "Germany", Price: 2000.0, ReleaseDate: releaseDate}
	tablet := storage.Product{Name: "iPad 2019", Country: "Spain", Price: 500.0, ReleaseDate: releaseDate}

	if err := db.CreateProducts(phone, laptop, tablet); err != nil {
		panic(err)
	}

	fetchedPhone, err := db.Product("Galaxy S20")
	if err != nil {
		panic(err)
	}
	fmt.Printf("The fetched product has the following information: %+v\n", fetchedPhone)

	phone.Price *= 0.5
	if err := db.UpdateProduct(phone); err != nil {
		panic(err)
	}

	fetchedPhone, err = db.Product("Galaxy S20")
	if err != nil {
		panic(err)
	}

	fmt.Printf("The fetched product's new information after applying the discount : %+v\n", fetchedPhone)
}
