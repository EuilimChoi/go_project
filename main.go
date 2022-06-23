package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	totalPages := getPages()

	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Req ", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".cardOutline")

	searchCards.Each(func(i int, card *goquery.Selection) {
		// id, _ := card.Attr("class")
		// splitID := strings.Split(id, " ")[4][4:]
		// title := card.Find(".jcs-JobTitle").Text()
		// location := card.Find(".companyLocation").Text()
		// summery := card.Find(".job-snippet").Text()
		salary := card.Find(".attribute_snippet").Text()
		fmt.Println(salary)
	})
}

func extractJob(card *goquery.Selection) {
	// id, _ := card.Attr("class")
	// splitID := strings.Split(id, " ")[4][4:]
	// title := card.Find(".jcs-JobTitle").Text()
	// location := card.Find(".companyLocation").Text()
	salary := card.Find(".attribute_snippet").Text()
	// summery := card.Find(".job-snippet").Text()
	fmt.Println(salary)
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Stauts : ", res.StatusCode)
	}
}
