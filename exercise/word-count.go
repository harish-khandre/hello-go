package main

import "strings"

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, value := range words {
		m[value] += 1
	}

	return m
}
