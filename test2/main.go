package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

type Store struct {
	Items []Item `validate:"dive"`
}

type Item struct {
	Name  string `validate:"required"` // 名前は必須
	Price int    `validate:"lte=100"`  // 値段は100以下
}

func main() {
	form := Store{
		Items: []Item{
			{Name: "item1", Price: 99},
			{Name: "", Price: 100},
			{Name: "item3", Price: 101},
		},
	}

	err := validator.New().Struct(&form)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		for _, e := range errors {
			fmt.Println(e)
		}
	}
}
