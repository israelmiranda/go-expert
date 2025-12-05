package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	URL        = "http://localhost:8080/cotacao"
	timeout    = 300 * time.Millisecond
	outputFile = "cotacao.txt"
)

type QuotationResponse struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", URL, nil)
	if err != nil {
		log.Fatalf("create request error: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Fatalf("client timeout (300ms): %v", ctx.Err())
		}
		log.Fatalf("http request error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("http status fail: %s", res.Status)
	}

	var quotationRes QuotationResponse
	if err := json.NewDecoder(res.Body).Decode(&quotationRes); err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Fatalf("client timeout (300ms) error: %v", ctx.Err())
		}
		log.Fatalf("JSON decoder error: %v", err)
	}

	quotation := quotationRes.Bid
	content := fmt.Sprintf("Dólar: %s", quotation)
	if err := os.WriteFile(outputFile, []byte(content), 0644); err != nil {
		log.Fatalf("error saving the quotation to the file %s: %v", outputFile, err)
	}

	log.Printf("quotation save on **%s** with: **%s**", outputFile, content)
}
