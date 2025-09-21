package testutils

import (
	"context"
	"testing"

	"github.com/sijiaoh/go-godot-template/api_server/ent"
)

func AssertResponseCode(t *testing.T, respCode, expectedCode int) {
	if respCode != expectedCode {
		t.Fatalf("expected response code %d, got %d", expectedCode, respCode)
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
