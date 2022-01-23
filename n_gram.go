package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type NextWord struct {
	word string
	prob float32
}

type Next []*NextWord

func (n Next) Len() int {
	return len(n)
}

func (n Next) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n Next) Less(i, j int) bool {
	return n[i].prob > n[j].prob
}

func (w *NextWord) Print() {
	fmt.Printf("The next word is %s, probability is %f\n", w.word, w.prob)
}

type NGram struct {
	gram    int
	wordBag map[string][]*NextWord
}

func NewNGram(gram int) *NGram {
	if gram < 2 {
		panic(errors.New("gram must be larger than 2"))
	}

	return &NGram{
		gram:    gram,
		wordBag: make(map[string][]*NextWord),
	}
}

func (n *NGram) Init(content []string) {
	var totalGram []string
	var wordGram []string

	for _, c := range content {
		splitWords := []string{"<s>"}
		for _, w := range c {
			splitWords = append(splitWords, string(w))
		}
		splitWords = append(splitWords, "</s>")

		for i := 0; i < len(splitWords)-n.gram+1; i++ {
			totalGram = append(totalGram, strings.Join(splitWords[i:i+n.gram], ""))
			wordGram = append(wordGram, strings.Join(splitWords[i:i+n.gram-1], ""))
		}
	}

	totalMap := make(map[string]int)
	wordMap := make(map[string]int)
	for _, t := range totalGram {
		if m, ok := totalMap[t]; ok {
			totalMap[t] = m + 1
		} else {
			totalMap[t] = 1
		}
	}
	for _, w := range wordGram {
		if m, ok := wordMap[w]; ok {
			wordMap[w] = m + 1
		} else {
			wordMap[w] = 1
		}
	}

	for k, v := range totalMap {
		word := k[0 : len(k)/n.gram*(n.gram-1)]
		if _, ok := wordMap[word]; !ok {
			n.wordBag[word] = make([]*NextWord, 0)
		}

		nextWordProb := float32(v) / float32(wordMap[word])
		n.wordBag[word] = append(n.wordBag[word], &NextWord{
			word: k[len(k)/n.gram*(n.gram-1):],
			prob: nextWordProb,
		})
	}
}

func (n *NGram) Predict(word string) []*NextWord {
	next, ok := n.wordBag[word]
	if ok {
		sort.Sort(Next(next))
		return next
	}

	return nil
}
