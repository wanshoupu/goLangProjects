package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
)

func worker(visited map[string]int, queue *Queue, wg *sync.WaitGroup, limit int) {
	defer wg.Done()
	for len(visited) < limit {
		nextUrl := queue.Pop()
		if nextUrl == nil {
			break
		}
		log.Println(nextUrl)
		urls := fetch(nextUrl.(string))
		for _, urlCandidate := range urls {
			if _, ok := visited[urlCandidate]; ok {
				continue
			}
			visited[urlCandidate] += 1
			queue.Push(urlCandidate)
		}
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
			if string(name) == "a" || string(name) == "link" {
				href := extractHref(tokenizer)
				thisUrl := normalizeUrl(href, resp.Request.URL)
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
	if parsedUrl.Path == "" {
		parsedUrl.Path = "/"
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
