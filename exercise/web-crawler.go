package main

import (
	"fmt"
	"sync"
)

// Fetcher interface defines the contract for fetching URLs
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// SafeUrlCache is a thread-safe cache to track visited URLs
type SafeUrlCache struct {
	visited map[string]bool
	mux     sync.Mutex
}

// HasVisited checks if a URL has been visited and marks it as visited if not
func (c *SafeUrlCache) HasVisited(url string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()

	if c.visited[url] {
		return true
	}

	c.visited[url] = true
	return false
}

// Crawl uses fetcher to recursively crawl pages starting with url, to a maximum of depth
func Crawl(url string, depth int, fetcher Fetcher) {
	// Create a WaitGroup to track all goroutines
	var wg sync.WaitGroup

	// Create a cache to track visited URLs
	cache := &SafeUrlCache{
		visited: make(map[string]bool),
	}

	// Start the recursive crawling
	crawlRecursive(url, depth, fetcher, &wg, cache)

	// Wait for all goroutines to complete
	wg.Wait()
}

// crawlRecursive handles the actual crawling logic
func crawlRecursive(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup, cache *SafeUrlCache) {
	// Check depth limit
	if depth <= 0 {
		return
	}

	// Skip if already visited
	if cache.HasVisited(url) {
		return
	}

	// Fetch the URL contents
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	// Launch a new goroutine for each URL found
	for _, u := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			crawlRecursive(u, depth-1, fetcher, wg, cache)
		}(u)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

func webCrawler() {
	Crawl("https://golang.org/", 4, fetcher)
}
