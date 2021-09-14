package main

import "testing"

func TestCreateShortURL(t *testing.T) {
	url := "https://example.com"
	shorturl := createShortURL(url)

	if shorturl.URL != url {
		t.Error("Incorrect URL Value")
	}
}
