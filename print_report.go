package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Println("REPORT for", baseURL)
	fmt.Println("=============================")
	sortedPages := sortPages(pages)
	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", page.count, page.url)
	}
}

type page struct {
	url  string
	count int
}

func sortPages(pages map[string]int) []page {
	pagesSlice := make([]page, 0, len(pages))
	for url, count := range pages {
		pagesSlice = append(pagesSlice, page{url, count})
	}

	sort.Slice(pagesSlice, func(i, j int) bool {
		if pagesSlice[i].count != pagesSlice[j].count {
			return pagesSlice[i].count > pagesSlice[j].count
		}
		return pagesSlice[i].url < pagesSlice[j].url
	})

	return pagesSlice
}