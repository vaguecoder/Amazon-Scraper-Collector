package main


import (
	"fmt"
	"log"
	"time"
	"bytes"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/PuerkitoBio/goquery"
)

type Inner struct {
	Name			string	`json:"name,omitempty"`
	ImageURL		string	`json:"imageURL,omitempty"`
	Desc			string	`json:"description,omitempty"`
	Price			string	`json:"price,omitempty"`
	TotalReviews	int		`json:"totalReviews,omitempty"`
}

type Outer struct {
	URL				string	`json:"url,omitempty"`
	Product			Inner	`json:"product,omitempty"`
}

type StatusObject struct {
	InsertedID 		string	`json:"InsertedID,omitempty"`
	MatchedCount	int		`json:"MatchedCount,omitempty"`
    ModifiedCount	int		`json:"ModifiedCount,omitempty"`
}

func scraper(url string) Outer {
	client := &http.Client{
        Timeout: 30 * time.Second,
	}

    response, err := client.Get(url)
    if err != nil {
        log.Fatal("Failed at getting list's GET(), ", err)
    }
	defer response.Body.Close()

    document, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
		log.Fatal("Failed at getting list's goquery body, ", err)
	}

	inner := Inner{
		Name:			getName(document),
		ImageURL:		getImageURL(document),
		Desc:			getDesc(document),
		Price:			getPrice(document),
		TotalReviews:	getTotalReviews(document),
	}
	outer := Outer{
		URL: 		url,
		Product:	inner,
	}
	return outer
}

func getFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the REST-API-in-Go!" +
				"\nPlease do POST Request to the API for scrapping Amazon Product Details.")
}

func postFunc(writer http.ResponseWriter, request *http.Request){
	decoder := json.NewDecoder(request.Body)
	form_data := Outer{}
	err := decoder.Decode(&form_data)
    if err != nil {
        panic(err)
	}

	form_data = scraper(form_data.URL)
	
	product_details, err := json.Marshal(form_data)
	if err != nil {
		log.Fatal("json.Marshal failed due to the error:", err)
	}
	
	collector_url := "http://backend:8081/collector"
    requestObject, err := http.NewRequest("POST", collector_url, bytes.NewBuffer(product_details))
    requestObject.Header.Set("content-type", "application/json")

    client := &http.Client{}
    response, err := client.Do(requestObject)
    if err != nil {
        panic(err)
    }
    defer response.Body.Close()

	var status StatusObject
	_ = json.NewDecoder(response.Body).Decode(&status)

	if status.MatchedCount == 0 {
		fmt.Fprintf(writer, "For URL: %s\nProduct details scraped and stored in database with ID: %s\n", form_data.URL, status.InsertedID)
	} else {
		if status.ModifiedCount == 0 {
			fmt.Fprintf(writer, "For URL: %s\nProduct details already exists in Database, and they match.\n", form_data.URL)
		} else {
			fmt.Fprintf(writer, "For URL: %s\nProduct details already exists in Database, and are updated with latest.\n", form_data.URL)
		}
	}

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/scraper", getFunc).Methods("GET")
	router.HandleFunc("/scraper", postFunc).Methods("POST")
    log.Fatal(http.ListenAndServe(":8080", router))
}
