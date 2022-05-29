package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
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
	defer resp.Body.Close()
	return extractUrls(resp)
}

func extractUrls(resp *http.Response) []string {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	r := bytes.NewReader(body)
	doc, err := html.Parse(r)
	if err != nil {
		return nil
	}

	var result []string
	nodes := traverse(doc)
	for _, n := range nodes {
		if n.Data == "href" {
			result = append(result, n.FirstChild.Data)
		}
	}
	return result
}

func traverse(root *html.Node) []html.Node {
	if root == nil {
		return []html.Node{}
	}
	var result []html.Node

	if root.Type == html.ElementNode {
		result = append(result, traverseSubtree(root)...)
	}
	for ; root.NextSibling != nil; root = root.NextSibling {
		result = append(result, traverse(root.NextSibling)...)
	}
	return result
}

func traverseSubtree(root *html.Node) []html.Node {
	var result []html.Node
	if root.Type != html.ElementNode {
		return result
	}
	if root.FirstChild != nil {

	}
}
