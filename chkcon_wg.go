package goutils

import (
	"fmt"
	"sync"
)

var cnt int
var wg sync.WaitGroup
var mut sync.Mutex

func CheckConnectivity(urls []string) int {
	for _, url := range urls {
		go checkUrl(url)
		wg.Add(1)
	}
	wg.Wait()
	return cnt
}

func checkUrl(url string) {
	defer wg.Done()
	_, err := GetWebSite(url)
	if err != nil {
		fmt.Println(url, "is down")
	} else {
		mut.Lock()
		defer mut.Unlock()
		fmt.Println(url, "is up")
		cnt += 1
	}
}
