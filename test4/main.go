package main

import (
	"fmt"
	"log"

	"github.com/go-playground/locales/ja_JP"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

type Item struct {
	Name string `json:"name" validate:"required"`
}

func main() {
	item := Item{Name: ""}
	v, trans := setupValidator()
	err := v.Struct(&item)
	if err != nil {
		errors := err.(validator.ValidationErrors).Translate(trans)
		for _, e := range errors {
			fmt.Println(e)
		}
	}
}

func setupValidator() (*validator.Validate, ut.Translator) {
	uTrans := ut.New(ja_JP.New(), ja_JP.New())

	// 変換処理生成
	jaTrans, found := uTrans.GetTranslator("ja_JP")
	if !found {
		log.Fatal("translator not found")
		return nil, nil
	}

	v := validator.New()
	// フィールド名を登録
	_ = jaTrans.Add("Name", "名前", false)

	// タグに対応するエラーメッセージを登録
	v.RegisterTranslation("required", jaTrans,
		func(ut ut.Translator) error {
			return ut.Add("required", "{0}は必須です。", false)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			// フィールド名取得
			fld, _ := ut.T(fe.Field())
			// エラーメッセージ取得
			t, _ := ut.T(fe.Tag(), fld)
			return t
		})

	return v, jaTrans
}
