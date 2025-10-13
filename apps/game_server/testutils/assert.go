package testutils

import (
	"testing"

	"github.com/sijiaoh/go-godot-template/game_server/utils"
)

func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func AssertWithError(t *testing.T, err error) {
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func AssertEqual[T comparable](t *testing.T, got, expected T) {
	if got != expected {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, expected T) {
	if got == expected {
		t.Fatalf("expected not %v, got %v", expected, got)
	}
}

func AssertStrPtrEqual(t *testing.T, got, expected *string) {
	if !utils.StrPtrEq(got, expected) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}

func AssertResponseCode(t *testing.T, respCode, expectedCode int) {
	if respCode != expectedCode {
		t.Fatalf("expected response code %d, got %d", expectedCode, respCode)
	}
}
