package main

import (
	"flag"
)

var word = flag.String("word", "中国", "The word needs to predict.")

func main() {
	flag.Parse()
	lines := loadSoHuData("data/Sohu.dat")
	ngram := NewNGram(3)
	ngram.Init(lines)

	ret := ngram.Predict(*word)
	for _, next := range ret {
		next.Print()
	}
}
