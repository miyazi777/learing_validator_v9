package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

type Item struct {
	Name string `json:"name" validate:"required,checkName"`
}

func checkNameValidate(fl validator.FieldLevel) bool {
	return fl.Field().String() == "Mike"
}

func main() {
	item := Item{Name: "Mike"}
	valid := validator.New()
	valid.RegisterValidation("checkName", checkNameValidate)
	err := valid.Struct(&item)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		for _, e := range errors {
			fmt.Println(e)
		}
	}
}
