package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat://somesitexyz.xyz" {
		return false
	}

	return true
}

func TestCheckWebister(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://codingtaco.com",
		"waat://somesitexyz.xyz",
	}

	want := map[string]bool{
		"https://google.com":     true,
		"https://codingtaco.com": true,
		"waat://somesitexyz.xyz": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
