package testutils

import (
	"context"
	"testing"

	"github.com/sijiaoh/go-godot-template/api_server/ent"
	"github.com/sijiaoh/go-godot-template/api_server/utils"
)

func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func AssertEqual(t *testing.T, got, expected interface{}) {
	if got != expected {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}

func AssertStrPtrEqual(t *testing.T, got, expected *string) {
	if !utils.StrPtrEq(got, expected) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}

func AssertRecordCount(t *testing.T, userQuery *ent.UserQuery, ctx context.Context, expectedCount int) {
	count, err := userQuery.Count(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if count != expectedCount {
		t.Fatalf("expected record count %d, got %d", expectedCount, count)
	}
}

func AssertResponseCode(t *testing.T, respCode, expectedCode int) {
	if respCode != expectedCode {
		t.Fatalf("expected response code %d, got %d", expectedCode, respCode)
	}
}
