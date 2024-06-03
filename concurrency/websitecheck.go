package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChan := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChan <- result{u, wc(u)}
		}(url)
	}

	for range urls {
		r := <-resultChan
		results[r.string] = r.bool
	}

	return results
}
