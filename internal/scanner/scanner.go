package scanner

import (
	mutatetion "RepoHound/internal/mutate"
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type Scanner struct {
	uList   []string
	workers int
	tr      *http.Transport
}

func New(workers int, f, proxy, keyword string) (Scanner, error) {

	var scanner = Scanner{}
	var tr = http.Transport{}
	var uList = make([]string, 0, 500)

	if proxy != "" {
		// parse proxy
		proxyUrl, err := url.Parse(proxy)
		if err != nil {
			return scanner, err
		}

		tr = http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	}

	file, err := os.Open(f)
	if err != nil {
		return scanner, err
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	for s.Scan() {
		uList = append(uList, s.Text())
	}

	if err := s.Err(); err != nil {
		return scanner, fmt.Errorf("error reading file: %w", err)
	}

	if keyword != "" {
		uList = mutatetion.Mutate(uList, keyword)
	}

	return Scanner{
		uList:   uList,
		workers: workers,
		tr:      &tr,
	}, nil

}
