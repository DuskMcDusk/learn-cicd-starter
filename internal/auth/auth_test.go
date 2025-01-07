package auth

import (
	"net/http"
	"testing"
)

func TestAuthError(t *testing.T) {
	header := http.Header{}
	_, err := GetAPIKey(header)
	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected %v, got %v", ErrNoAuthHeaderIncluded, err)
	}
}

func TestAuthKey(t *testing.T) {
	header := http.Header{}
	expected := "TestKey"
	header.Add("Authorization", "ApiKey "+expected)
	got, err := GetAPIKey(header)
	if err != nil {
		t.Fatal("Unexpected error")
	}
	if got != expected {
		t.Fatalf("expected %v, got %v", got, expected)
	}
}
