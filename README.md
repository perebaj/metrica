# metrica

What we are trying to solve here? A simple race condition problem, here multiple pieces of code are trying to access the same shared data!


# Api Example

Request
```bash
    curl -i -X 'GET' http://localhost:8080/count
```


## Resources

This example is really close with the problem that we have here

[Mutex](https://golangbot.com/mutex/)
[Atomic Counters](https://gobyexample.com/atomic-counters)
[Bjorn Rabenstein - Prometheus: Designing and Implementing a Modern Monitoring Solution in Go](https://www.youtube.com/watch?v=1V7eJ0jN8-E)
