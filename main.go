package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
    "time"
)

type quoteResp struct {
    Content string `json:"content"`
    Author  string `json:"author"`
}

func fetchQuote() (*quoteResp, error) {
    client := &http.Client{Timeout: 5 * time.Second}
    resp, err := client.Get("https://api.quotable.io/random")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected status: %s", resp.Status)
    }
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    var q quoteResp
    if err := json.Unmarshal(body, &q); err != nil {
        return nil, err
    }
    return &q, nil
}

func main() {
    q, err := fetchQuote()
    if err != nil {
        fmt.Fprintln(os.Stderr, "error fetching quote:", err)
        os.Exit(1)
    }
    fmt.Printf("\"%s\" – %s\n", q.Content, q.Author)
}
