package main

import (
        "fmt"
        "github.com/gocolly/colly"
)


// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
        // Instantiate default collector
        c := colly.NewCollector()

        // 3
        c.OnHTML("a[href]", func(e *colly.HTMLElement) {
                // Find link using an attribute selector
                // Matches any element that includes href=""
                link := e.Attr("href")

                // Print link
                fmt.Printf("Link found: %q -> %s\n", e.Text, link)

                // Visit link
                e.Request.Visit(link)
        })

        // 1
        c.OnRequest(func(r *colly.Request) {
                fmt.Println("Visiting", r.URL)
        })

        // alternative #3 if something goes wrong 
        c.OnError(func(_ *colly.Response, err error) {
                fmt.Println("Something went wrong:", err)
        })

        // 2
        c.OnResponse(func(r *colly.Response) {
                fmt.Println("Visited", r.Request.URL)
        })

        // 4
        c.OnScraped(func(r *colly.Response) {
                fmt.Println("Finished", r.Request.URL)
        })

        
        // Start scraping on https://hackerspaces.org
        c.Visit("https://hackerspaces.org/")
}
