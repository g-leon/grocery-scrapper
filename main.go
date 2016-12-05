package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"github.com/PuerkitoBio/goquery"
	"github.com/g-leon/grocery-scrapper/format"
	"github.com/g-leon/grocery-scrapper/product"
	"github.com/g-leon/grocery-scrapper/scraper"
)

const URL = "http://hiring-tests.s3-website-eu-west-1.amazonaws.com/2015_Developer_Scrape/5_products.html"

func main() {
	products, total := getProducts(URL)
	m := map[string]interface{}{
		"results": products,
		"total":   strconv.FormatFloat(total, 'f', 2, 64),
	}
	res, err := json.Marshal(m)
	res = format.Encoding(res)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(res))
	}
}

// getProducts returns a list of products and the total sum of their prices
func getProducts(url string) ([]*product.Product, float64) {
	products := make([]*product.Product, 0)
	var total float64

	s := scraper.New(url)

	// Select the entire page
	pSelector := s.Find(".product")

	pChan := make(chan *product.Product, pSelector.Size())

	scrapeProducts(pChan, pSelector)
	close(pChan)

	total = 0
	for p := range pChan {
		products = append(products, p)
		total += p.UnitPrice
	}

	return products, total
}

// scrapeProducts will fetch each product item from the web page identified by the given URL
func scrapeProducts(pChan chan *product.Product, pSelector *goquery.Selection) {
	var wg sync.WaitGroup
	wg.Add(pSelector.Size())
	pSelector.Each(func(i int, s *goquery.Selection) {
		title := s.Find("h3")
		url, _ := title.Find("a").Attr("href")
		price := s.Find(".pricePerUnit").Text()
		go func(pChan chan *product.Product) {
			defer wg.Done()
			pChan <- product.New(title.Text(), price, url)

		}(pChan)
	})
	wg.Wait()
}