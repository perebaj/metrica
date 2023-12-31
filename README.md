# metrica

What we are trying to solve here? A simple race condition problem, where multiple pieces of code are trying to access the same shared data!

![image](assets/metrica.png)


## Get starting

All commands could be accessed typing: `make help`

To test the code using containers🐋

`make dev/test` & `make dev/lint`


## LoadTest 

The load tests it's using the following rules:

	- 100 concurrent users
	- 10000 requests in total

![image](assets/loadtest.png)

## Resources

Some resources that were useful to solve it

- [Mutex](https://golangbot.com/mutex/)
- [Atomic Counters](https://gobyexample.com/atomic-counters)
- [Bjorn Rabenstein - Prometheus: Designing and Implementing a Modern Monitoring Solution in Go](https://www.youtube.com/watch?v=1V7eJ0jN8-E)
- [go file mutex](https://echorand.me/posts/go-file-mutex/)
