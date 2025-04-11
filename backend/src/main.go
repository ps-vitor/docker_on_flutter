package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type ScrapeResult struct {
	URL  string `json:"url"`
	Text string `json:"text"`
	Err  string `json:"err,omitempty"`
}

type Job struct {
	URL      string
	Interval time.Duration
}

func scrapePage(url string, ch chan<- ScrapeResult) {
	customClient := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := customClient.Get(url)
	if err != nil {
		ch <- ScrapeResult{URL: url, Text: "", Err: fmt.Sprintf("Request error: %v", err)}
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		ch <- ScrapeResult{URL: url, Text: "", Err: fmt.Sprintf("Parse error: %v", err)}
		return
	}

	selection := doc.Find("div.content.clearfix")
	text := selection.Text()
	if text == "" {
		text = "Not found."
	}

	ch <- ScrapeResult{URL: url, Text: text}
}

func scrap(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "Missing URL", http.StatusBadRequest)
	}

	ch := make(chan ScrapeResult)
	go scrapePage(url, ch)
	result := <-ch

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

var jobs []Job

func startScraper(job Job) {
	ticker := time.NewTicker(job.Interval)
	go func() {
		for range ticker.C {
			ch := make(chan ScrapeResult)
			go scrapePage(job.URL, ch)
			result := <-ch
			fmt.Printf("Scraped [%s]: %.30s...\n", result.URL, result.Text)
			break
		}
	}()
}

func addJobHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	intervalStr := r.URL.Query().Get("interval")

	intervalSec, err := strconv.Atoi(intervalStr)
	if err != nil || url == "" {
		http.Error(w, "Invalid URL or interval", http.StatusBadRequest)
		return
	}

	job := Job{
		URL:      url,
		Interval: time.Duration(intervalSec) * time.Second,
	}
	jobs = append(jobs, job)
	startScraper(job)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("Job for %s started every %d seconds", url, intervalSec),
	})
}

func main() {
	http.HandleFunc("/add-job", addJobHandler)
	http.HandleFunc("/scrap", scrap)
	fmt.Println("Server running on http://0.0.0.0:8080/scrape?url=")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
