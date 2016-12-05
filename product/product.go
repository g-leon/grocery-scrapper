package product

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/g-leon/grocery-scrapper/numbers"
	"github.com/g-leon/grocery-scrapper/scraper"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Product struct {
	url         string
	Title       string  `json:"title"`
	Size        string  `json:"size"`
	UnitPrice   float64 `json:"unit_price"`
	Description string  `json:"description"`
}

func New(title string, price string, url string) *Product {
	p := &Product{}
	p.url = strings.TrimSpace(url)
	p.Title = strings.TrimSpace(title)
	pPage := p.getPageScraper()
	p.Size = pPage.Size()
	p.Description = strings.TrimSpace(p.description(pPage))
	p.UnitPrice = p.unitPrice(price)
	return p
}

// unitPrice returns the price found by the selector on its working page
func (p *Product) unitPrice(price string) float64 {
	r := regexp.MustCompile(`[-+]?\d*\.\d+|\d+`)
	sp := r.FindStringSubmatch(strings.TrimSpace(price))

	if len(sp) != 0 {
		fp, err := strconv.ParseFloat(sp[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		return numbers.RoundFloat(float64(fp), .5, 2)
	}
	return 0
}

// description returns the description found by the selector on its working page
func (p *Product) description(pPage *scraper.Scraper) string {
	d := ""
	pPage.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("name"); name == "description" {
			d, _ = s.Attr("content")
		}
	})
	return d
}

// getPageScraper returns a scraper that works on the specified URL
func (p *Product) getPageScraper() *scraper.Scraper {
	return scraper.New(p.url)
}

// MarshalJSON is a custom implementation of the MarshalJSON interface
// and it makes sure that the unit price will have two decimals precision
// when a Product
func (p *Product) MarshalJSON() ([]byte, error) {
	type Alias Product
	return json.Marshal(&struct {
		UnitPrice string `json:"unit_price"`
		*Alias
	}{
		UnitPrice: strconv.FormatFloat(p.UnitPrice, 'f', 2, 64),
		Alias:     (*Alias)(p),
	})
}
