package main

import "testing"


func TestGetProducts(t *testing.T) {
	var url = "http://hiring-tests.s3-website-eu-west-1.amazonaws.com/2015_Developer_Scrape/5_products.html"
	products, total := getProducts(url)

	if total == 0 {
		t.Error("Total is 0")
	}
	if len(products) == 0 {
		t.Error("Number of products is 0")
	}
}
