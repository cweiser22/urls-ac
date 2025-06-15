package models

import "testing"

func TestURL_Equals(t *testing.T) {
	url1 := &URLMapping{LongURL: "http://example.com"}
	url2 := &URLMapping{LongURL: "http://example.com"}

	if !url1.Equals(url2) {
		t.Error("URLs should be equal")
	}

	url3 := &URLMapping{LongURL: "http://youtube.com"}
	if url1.Equals(url3) {
		t.Error("URLs should not be equal")
	}
}
