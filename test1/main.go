package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

type Item struct {
	Item1  string `validate:"required"`        // 必須チェック(空文字でもエラー)
	Item2  string `validate:"len=10"`          // 長さチェック(11文字以上でエラー)
	Item3  string `validate:"min=3,max=5"`     // 長さチェック(2文字以下でエラー、6文字以上でエラー)
	Item4  string `validate:"eq=abc"`          // 値チェック(abc以外はエラー)
	Item5  string `validate:"ne=abc"`          // 値チェック(abcはエラー)
	Item6  string `validate:"oneof=red green"` // 配列に設定されているもの以外エラー(red green以外エラー)
	Item7  int    `validate:"gt=10"`           // 数値チェック(n > 10かチェック)
	Item8  int    `validate:"gte=10"`          // 数値チェック(n >= 10かチェック)
	Item9  int    `validate:"lt=10"`           // 数値チェック(n < 10かチェック)
	Item10 int    `validate:"lte=10"`          // 数値チェック(n <= 10かチェック)
	Item11 string `validate:"eqfield=Item12"`  // 指定した他フィールドと同じか？
	Item12 string `validate:""`                //
	Item13 string `validate:"nefield=Item14"`  // 指定した他フィールドと違うか？
	Item14 string `validate:""`                //
	Item15 string `validate:"alpha"`           // 文字のみか？
	Item16 string `validate:"alphanum"`        // 文字と数値のみか？
	Item17 string `validate:"numeric"`         // 数値のみか？
	Item18 string `validate:"email"`           // メール形式か？
}

func main() {
	form := Item{
		Item1:  "",
		Item2:  "12345678901",
		Item3:  "123456",
		Item4:  "abcd",
		Item5:  "abc",
		Item6:  "blue",
		Item7:  10,
		Item8:  9,
		Item9:  10,
		Item10: 11,
		Item11: "value1",
		Item12: "value2",
		Item13: "value1",
		Item14: "value1",
		Item15: "abc1",
		Item16: "123x,",
		Item17: "123x",
		Item18: "test.tset.com",
	}

	fmt.Println(form)
	fmt.Println("------")
	err := validator.New().Struct(&form)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		for _, e := range errors {
			fmt.Println(e)
		}
	}
}
