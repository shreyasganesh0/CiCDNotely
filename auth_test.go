package main

import (
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	authheader := make(http.Header)
	canon_s := http.CanonicalHeaderKey("Autorization")

	var short_val string = "a"
	authheader.Set(canon_s, short_val)

	_, got_err := auth.GetAPIKey(authheader)

	if reflect.DeepEqual(nil, got_err) {
		t.Fatal("Expected nil for error got\n")
	}
}
