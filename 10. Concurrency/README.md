# Concurrency - Goroutines

A Goroutine is a function that is capable of running concurrently with
other functions. To create a goroutine we use a keyword `go` followed
by a function invocation.

- `goroutines`, the basic unit of concurrency in Go, which let us check
more than one website at the same time.
- `anonymous functions`, which we used to start each of the concurrent processes that check websites.
- `channels`, to help organize and control the communication between the different processes, allowing us to avoid a race condition bug.
- `race detector`, which helped us debug problems with concurrent code.