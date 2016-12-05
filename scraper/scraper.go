package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
)

type Scraper struct {
	doc *goquery.Document
	url string
}

func New(url string) *Scraper {
	s := new(Scraper)
	s.url = url
	s.doc = s.buildDocument()
	return s
}

// Size returns the size of the document in kb
func (s *Scraper) Size() string {
	size := float64(s.contentSize()) / 1024
	return strconv.FormatFloat(size, 'f', 2, 64) + "kb"
}

// Find is a wrapper for *goquery.Document().Find()
func (s *Scraper) Find(selector string) *goquery.Selection {
	return s.doc.Find(selector)
}

// contentSize returns the size of the web page
func (s *Scraper) contentSize() int64 {
	webPage, err := s.doc.Html()
	if err != nil {
		log.Fatal(err)
	}
	return int64(len(webPage))
}

// buildDocument returns a *goquery.Document built from a http response body
func (s *Scraper) buildDocument() *goquery.Document {
	resp, err := http.Get(s.url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}
