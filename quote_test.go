package main

import "testing"

func TestFetchQuote(t *testing.T) {
    q, err := fetchQuote()
    if err != nil {
        t.Fatalf("fetchQuote returned error: %v", err)
    }
    if q.Content == "" || q.Author == "" {
        t.Fatalf("quote fields should not be empty: %+v", q)
    }
}
