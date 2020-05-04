package main

import (
	"fmt"
	"strings"
)

type printCondition func(w string, c int) bool
type wordCountMap map[string]int

func main() {
	txt := `As far as eye could reach he saw nothing 
but the stems of the great plants about him receding 
in the violet shade, and far overhead the multiple 
transparency of huge leaves filtering the sunshine to 
the solemn splendour of twilight in which he walked. 
Whenever he felt able he ran again; the ground continued 
soft and springy, covered with the same resilient weed 
which was the first thing his hands had touched in 
Malacandra. Once or twice a small red creature scuttled 
across his path, but otherwise there seemed to be no 
life stirring in the wood; nothing to fear—except the 
fact of wandering unprovisioned and alone in a forest 
of unknown vegetation thousands or millions of miles 
beyond the reach or knowledge of man.`

	wordCount(txt).print(func(w string, c int) bool {
		return c > 2
	})
}

func wordCount(txt string) wordCountMap{
	m := wordCountMap{}
	words := strings.Fields(strings.ToLower(txt))
	for _, word := range words {
		i := strings.Trim(word, ".;-,")
		m[i]++
	}

	return m
}

func (m wordCountMap) print(condition printCondition) {
	for word, count := range m {
		if !condition(word, count) {
			continue
		}
		fmt.Println(word, count)
	}
}
