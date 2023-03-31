package main

func Search(dictionary map[string]string, word string) string {
	val, ok := dictionary[word]
	if ok {
		return val
	}
	return ""
}
