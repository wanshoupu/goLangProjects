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
	visited := make(map[string]int)
	queue := NewQueue(1000)

	for _, pa := range pas {
		queue.Push(pa)
	}

	for len(visited) < 1000 && valueSum(visited) < 1000 {
		pa := queue.Pop()
		if pa == nil {
			break
		}
		fmt.Println(pa)
		visited[pa.(string)] += 1
		if _, ok := visited[pa.(string)]; !ok {
			continue
		}
		urls := fetch(pa.(string))
		for _, u := range urls {
			queue.Push(u)
		}
	}

	fmt.Println("done")
	for k, v := range visited {
		fmt.Println(k, v)
	}
}

func valueSum(myDict map[string]int) int {
	result := 0
	for _, v := range myDict {
		result += v
	}
	return result
}

func fetch(pa string) []string {
	client := &http.Client{}
	resp, err := client.Get(pa)
	if err != nil {
		fmt.Println(err)
		return []string{}
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
	} else if parsedUrl.Scheme == "" {
		parsedUrl.Scheme = "http"
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
