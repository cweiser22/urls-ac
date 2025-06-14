package models

import "testing"

func TestURL_Equals(t *testing.T) {
	url1 := &URL{LongURL: "http://example.com"}
	url2 := &URL{LongURL: "http://example.com"}

	if !url1.Equals(url2) {
		t.Error("URLs should be equal")
	}

	url3 := &URL{LongURL: "http://youtube.com"}
	if url1.Equals(url3) {
		t.Error("URLs should not be equal")
	}
}
