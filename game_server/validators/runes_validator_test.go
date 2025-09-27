package validators_test

import (
	"testing"

	"github.com/sijiaoh/go-godot-template/game_server/validators"
)

type RunesTestStruct struct {
	Name string `validate:"min_runes=3,max_runes=5"`
}

func TestMinRunes(t *testing.T) {
	rts := RunesTestStruct{Name: "您好"}
	err := validators.Validate().Struct(rts)
	if err == nil {
		t.Fatal("UTF-8 Rune字数小于限制值，但验证通过了")
	}

	rts.Name = "您好啊"
	err = validators.Validate().Struct(rts)
	if err != nil {
		t.Fatal("UTF-8 Rune字数在限制值范围内，但验证未通过")
	}
}

func TestMaxRunes(t *testing.T) {
	rts := RunesTestStruct{Name: "您好啊，欢迎"}
	err := validators.Validate().Struct(rts)
	if err == nil {
		t.Fatal("UTF-8 Rune字数大于限制值，但验证通过了")
	}

	rts.Name = "您好啊欢迎"
	err = validators.Validate().Struct(rts)
	if err != nil {
		t.Fatal("UTF-8 Rune字数在限制值范围内，但验证未通过")
	}
}
