package main

import "testing"

const URL = "http://hiring-tests.s3-website-eu-west-1.amazonaws.com/2015_Developer_Scrape/5_products.html"

func TestGetProducts(t *testing.T) {
	products, total := getProducts(URL)

	if total == 0 {
		t.Error("Total is 0")
	}
	if len(products) == 0 {
		t.Error("Number of products is 0")
	}
}
