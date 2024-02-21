package main

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	res := make(map[string]bool)
	for _, url := range urls {
		go func() { //anonymous function
			res[url] = wc(url)
		}() //execute immediately upon declaration
	}
	return res
}
