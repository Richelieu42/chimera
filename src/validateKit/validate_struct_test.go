package validateKit

import (
	"fmt"
	"testing"
)

func TestStructExcept(t *testing.T) {
	type inner struct {
		Sex string `validate:"required"`
	}
	type bean struct {
		Id    int    `validate:"required"`
		Name  string `validate:"required"`
		Inner *inner
	}

	b := &bean{
		Id:   1,
		Name: "",
		Inner: &inner{
			Sex: "",
		},
	}

	if err := Struct(b); err != nil {
		fmt.Println(err)
		/*
			Key: 'bean.Name' Error:Field validation for 'Name' failed on the 'required' tag
			Key: 'bean.Inner.Sex' Error:Field validation for 'Sex' failed on the 'required' tag
		*/
	} else {
		fmt.Println("ok")
	}

	if err := StructExcept(b, "Name", "Inner.Sex"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok") // ok
	}
}

func TestStructPartial(t *testing.T) {
	type inner struct {
		Sex string `validate:"required"`
	}
	type bean struct {
		Id    int    `validate:"required"`
		Name  string `validate:"required"`
		Inner *inner
	}

	b := &bean{
		Id:   1,
		Name: "",
		Inner: &inner{
			Sex: "",
		},
	}

	if err := StructPartial(b, "Inner.Sex"); err != nil {
		fmt.Println(err) // Key: 'bean.Inner.Sex' Error:Field validation for 'Sex' failed on the 'required' tag
	} else {
		fmt.Println("ok")
	}
}
