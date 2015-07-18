package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"golang.org/x/net/html"
)

var urls = []string{
	"http://codingarena.in",
	"http://9piecesof8.com",
}

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

func asyncHttpGets(urls []string) []*HttpResponse {
	ch := make(chan *HttpResponse)
	responses := []*HttpResponse{}
	for _, url := range urls {
		go func(url string) {
			resp, err := http.Get(url)
			doc, err := html.Parse(resp.Body)
			parseHTML(doc)
			ch <- &HttpResponse{url, resp, err}
		}(url)
	}

	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			responses = append(responses, r)
			if len(responses) == len(urls) {
				return responses
			}
		case <-time.After(100 * time.Millisecond):
			fmt.Printf(".")
		}
	}
}

func parseHTML(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, element := range n.Attr {
			if element.Key == "href" && element.Val != "#" {
				fmt.Printf("LINK: %s\n", element.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseHTML(c)
	}
}

func main() {
	runtime.GOMAXPROCS(8)
	results := asyncHttpGets(urls)
	for _, result := range results {
		fmt.Printf("%s status: %s\n", result.url,
			result.response.Status)
	}
}
