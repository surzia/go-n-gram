package main

func main() {
	lines := loadSoHuData("data/Sohu.dat")
	ngram := NewNGram(3)
	ngram.Init(lines)

	ret := ngram.Predict("中国")
	for _, next := range ret {
		next.Print()
	}
}
