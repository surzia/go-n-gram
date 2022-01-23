# README

This is an N-gram model implemented by go.

## N-gram
A language model is a probability distribution over sequences of words `W`, namely:

![eq1](img/eq1.png)

According to the chain rule,

![eq2](img/eq2.png)

We can use [Maximum Likelihood Estimation](https://leimao.github.io/blog/Maximum-Likelihood-Estimation-Ngram/)

![eq3](img/eq3.png)

## Usage
You can run it directly, as:
```shell
go run . -word "中国"
```
or run it after compilation:
```shell
go build .
./go-n-gram -word "中国"
```
the output is:
```
The next word is 人, probability is 0.071429
The next word is 扶, probability is 0.055556
The next word is 的, probability is 0.039683
The next word is 社, probability is 0.039683
The next word is ，, probability is 0.039683
......
```