# Word Typing Prediction

> What's the next word?

## N-gram
A language model is a probability distribution over sequences of words, namely:

![](https://latex.codecogs.com/svg.image?p(w_{1},w_{2},...,w_{n}))

According to the chain rule,

![](https://latex.codecogs.com/svg.image?p(w_{1},w_{2},...,w_{n})=p(w_{1})p(w_{2}|w_{1})\cdot&space;\cdot&space;\cdot&space;p(w_{n}|w_{n-1},w_{n-2},...,w_{1}))

When we calculate the word 'python', we can use [Maximum Likelihood Estimation](https://leimao.github.io/blog/Maximum-Likelihood-Estimation-Ngram/)

![](https://latex.codecogs.com/svg.image?p(n|pytho)=\frac{count(python)}{count(pytho)})