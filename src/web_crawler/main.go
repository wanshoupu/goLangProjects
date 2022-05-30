package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

/**
The web crawler will utilize multiple threads.
It will be able to crawl all the particular web pages of a website.
It will be able to report back any 2XX and 4XX links.
It will take in the domain name from the command line.
It will avoid the cyclic traversal of links.

*/
func main() {
	pas := os.Args
	for i, pa := range pas {
		fmt.Println(i, pa)
		body := fetch(pa)
		fmt.Println(body)

	}
}

func fetch(pa string) []string {
	client := &http.Client{}
	resp, err := client.Get(pa)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	return extractUrls(resp)
}

func extractUrls(resp *http.Response) []string {
	tokenizer := html.NewTokenizer(resp.Body)
	var hrefMap = make(map[string]int)
	for {
		next := tokenizer.Next()
		if next == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				//end of the file, break out of the loop
				break
			}
			//otherwise, there was an error tokenizing,
			//which likely means the HTML was malformed.
			log.Fatalf("error tokenizing HTML: %v", tokenizer.Err())
		} else if next == html.StartTagToken {
			name, _ := tokenizer.TagName()
			if string(name) == "a" {
				href := extractHref(tokenizer)
				thisUrl := normalizeUrl(href, resp.Request.URL)
				log.Println(thisUrl)
				hrefMap[thisUrl] += 1
			}
		}
	}
	var result = make([]string, len(hrefMap))
	i := 0
	for k := range hrefMap {
		result[i] = k
		i++
	}
	return result
}

func normalizeUrl(thisUrl string, req *url.URL) string {
	parsedUrl, err := url.Parse(thisUrl)
	if err != nil {
		return ""
	}
	if parsedUrl.Host == "" {
		parsedUrl.Host = req.Host
		parsedUrl.Scheme = req.Scheme
	}
	return parsedUrl.String()
}

func extractHref(tokenizer *html.Tokenizer) string {
	for {
		key, val, hasMore := tokenizer.TagAttr()
		if string(key) == "href" {
			return string(val)
		}
		if !hasMore {
			break
		}
	}
	return ""
}
