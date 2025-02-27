package scanner

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type Result struct {
	Username string
	Repos    []string
}

func (s *Scanner) Run() []Result {

	var wg sync.WaitGroup
	var rg sync.WaitGroup
	var wq = make(chan string)
	var wr = make(chan Result)
	var results []Result

	// consume the result
	go func() {
		for r := range wr {
			results = append(results, r)
		}
		rg.Done()
	}()
	rg.Add(1)

	// start workers
	for i := 0; i < s.workers; i++ {
		wg.Add(1)
		go startWorker(&wg, wq, wr, s.tr)
	}

	for _, u := range s.uList {
		wq <- u
	}

	close(wq)
	wg.Wait()

	close(wr)
	rg.Wait()

	return results

}

func startWorker(wg *sync.WaitGroup, wq <-chan string, wr chan<- Result, tr *http.Transport) {

	for u := range wq {
		// endpoint := fmt.Sprintf("https://api.github.com/users/%s/repos", u)
		endpoint := fmt.Sprintf("https://github.com/%s?tab=repositories", u)

		links, err := doReq(endpoint, tr)
		if err != nil {
			continue
		}

		wr <- Result{Username: u, Repos: links}

	}

	wg.Done()
}

func doReq(endpoint string, tr *http.Transport) ([]string, error) {
	var links []string

	client := &http.Client{Transport: tr}

	res, err := client.Get(endpoint)
	if err != nil {
		return links, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return links, err
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return links, err
	}

	// Find all <a> tags with itemprop="name codeRepository"
	doc.Find(`a[itemprop="name codeRepository"]`).Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			links = append(links, href)
		}
	})

	return links, nil
}
