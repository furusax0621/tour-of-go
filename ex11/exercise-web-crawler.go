// Exercise: Web Crawler
// この演習では、ウェブクローラ( web crawler )を並列化するため、Goの並行性の特徴を使います。
// 同じURLを2度取ってくることなく並列してURLを取ってくるように、 Crawl 関数を修正してみてください。
package main

import (
	"fmt"
	"sync"
	"time"
)

// CrawlChecker はFetchしたURLのリストと数を管理する構造体です。
type CrawlChecker struct {
	URLs  map[string]int
	count int
	mux   sync.Mutex
}

// AddCount はFetch済みURLの数をカウントアップします。
func (c *CrawlChecker) AddCount() {
	c.mux.Lock()
	c.count++
	c.mux.Unlock()
}

// AddURL はFetchするURLのリストに url を追加します。
func (c *CrawlChecker) AddURL(url string) {
	c.mux.Lock()
	c.URLs[url] = 0
	c.mux.Unlock()
}

// Fetcher インターフェース。
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	// 同じURLを2回Fetchしないようにする
	_, ok := checker.URLs[url]
	if ok {
		return
	}

	// 新しいURLをリストに登録してFetch
	checker.AddURL(url)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		checker.AddCount()
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)
	}
	checker.AddCount()

	return
}

func main() {
	// すべての Web Crawl が完了したことを通知するためのチャネル
	quit := make(chan int)
	// timer は、Fetch予定のURLがすべてFetchしたか監視するための関数
	timer := func() {
		for checker.count < len(checker.URLs) {
			time.Sleep(10 * time.Millisecond)
		}
		quit <- 0
	}
	go timer()

	Crawl("https://golang.org/", 4, fetcher)
	<-quit
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

var checker = &CrawlChecker{URLs: make(map[string]int)}
