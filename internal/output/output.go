package output

import (
	"RepoHound/internal/scanner"
	"fmt"
)

func Print(results []scanner.Result, json bool) {

	if !json {
		fmt.Printf("Found The Following Repos:\n")
		for _, r := range results {
			for _, link := range r.Repos {
				fmt.Println(link)
			}
		}
	}
}
