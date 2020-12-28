package main

import (
	"regexp"
	"strings"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func getName(document *goquery.Document) string {
	var name string
	h1 := document.Find("h1#title").First()
	name = h1.Find("span#productTitle").Text()
	name = strings.TrimSpace(name)

	if name != "" {
		return name
	} else {
		return "Name Not Found!"
	}
}

func getImageURL(document *goquery.Document) string {
	var imageURL string
	document.Find("div#imgTagWrapperId").First().Each(func(i int, div *goquery.Selection) {
		str, _ := div.Find("img").Attr("data-a-dynamic-image")
		pattern, _:= regexp.Compile("https:\\/\\/.*?.jpg")
		img := pattern.FindAllString(str, -1)
		if len(img) > 1 {
			imageURL = img[len(img)-1]
		}
	})

	if imageURL != "" {
		return imageURL
	} else {
		return "Image URL Not Found!"
	}
}

func getDesc(document *goquery.Document) string {
	var description string
	document.Find("div#feature-bullets").First().Find("li").Each(func(i int, li *goquery.Selection) {
		if i != 0 {
			description += strings.TrimSpace(li.Find("span.a-list-item").Text()) + ". "
		}
	})

	if description != "" {
		return description
	} else {
		return "Description Not Found!"
	}
}

func getPrice(document *goquery.Document) string {
	var price string
	pattern, _:= regexp.Compile("(\\$[,0-9]*(\\.)([0-9])+)")

	price = document.Find("span#priceblock_ourprice").First().Text()

	if !pattern.Match([]byte(price)) {
		str := document.Find("ul.a-unordered-list").First().Find("li#edition_0").First().Find("span.a-size-mini").Text()
		temp := pattern.FindAllString(str, 1)
		if len(temp) > 1 {
			price = temp[0]
		}
	}

	if !pattern.Match([]byte(price)) {
		html, _ := document.Html()
		price = pattern.FindString(string(html))
	}

	if price != "" {
		return price
	} else {
		return "Price Not Found!"
	}
}

func getTotalReviews(document *goquery.Document) int {
	var totalReviews int
	temp_str := document.Find("span#acrCustomerReviewText").First().Text()
	temp_str = strings.ReplaceAll(temp_str, ",", "")
	temp_str = strings.Split(temp_str, " ")[0]
	totalReviews, _ = strconv.Atoi(temp_str)

	return totalReviews
}

