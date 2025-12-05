package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	quotationURL     = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	quotationTimeout = 200 * time.Millisecond
	dbTimeout        = 10 * time.Millisecond

	serverPort = ":8080"
)

type QuotationAPIResponse struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Bid        string `json:"bid"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

type QuotationResponse struct {
	Bid string `json:"bid"`
}

var db *sql.DB

func main() {
	if err := initDB(); err != nil {
		log.Fatalf("error starting DB: %v", err)
	}
	defer db.Close()

	http.HandleFunc("/cotacao", handleQuotation)
	log.Printf("server listening on port %s", serverPort)

	if err := http.ListenAndServe(serverPort, nil); err != nil {
		log.Fatal(err)
	}
}

func handleQuotation(w http.ResponseWriter, r *http.Request) {
	ctxAPI, cancelAPI := context.WithTimeout(r.Context(), quotationTimeout)
	defer cancelAPI()

	quotation, err := getQuotation(ctxAPI)
	if err != nil {
		if ctxAPI.Err() == context.DeadlineExceeded {
			log.Printf("quotation api timeout (200ms): %v", ctxAPI.Err())
			http.Error(w, "timeout limit to get a quotation", http.StatusRequestTimeout)
			return
		}
		http.Error(w, fmt.Sprintf("error obtaining a quotation: %v", err), http.StatusInternalServerError)
		return
	}

	ctxDB, cancelDB := context.WithTimeout(r.Context(), dbTimeout)
	defer cancelDB()

	if err := saveQuotation(ctxDB, quotation); err != nil {
		if ctxDB.Err() == context.DeadlineExceeded {
			log.Printf("save quotation DB timeout (10ms): %v", ctxAPI.Err())
		} else {
			log.Printf("DB error: %v", err)
		}
	} else {
		log.Printf("quotation save: %s", quotation.USDBRL.Bid)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := QuotationResponse{Bid: quotation.USDBRL.Bid}
	json.NewEncoder(w).Encode(response)
}

func getQuotation(ctx context.Context) (*QuotationAPIResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", quotationURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var quotation QuotationAPIResponse
	if err := json.Unmarshal(body, &quotation); err != nil {
		return nil, err
	}

	return &quotation, nil
}

func saveQuotation(ctx context.Context, quotation *QuotationAPIResponse) error {
	stmt, err := db.PrepareContext(ctx,
		"INSERT INTO quotes (code, codein, bid, create_date) VALUES (?, ?, ?, ?)",
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		quotation.USDBRL.Code, quotation.USDBRL.Codein, quotation.USDBRL.Bid, quotation.USDBRL.CreateDate,
	)
	if err != nil {
		return err
	}

	return nil
}

func initDB() error {
	var err error
	db, err = sql.Open("sqlite3", "quotes.db")
	if err != nil {
		return err
	}

	query := `
	CREATE TABLE IF NOT EXISTS quotes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code TEXT,
		codein TEXT,
		bid TEXT,
		create_date TEXT
	);`

	if _, err := db.Exec(query); err != nil {
		return err
	}

	return nil
}
