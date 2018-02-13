# Package bufpool

bufpool provides a thin wrapper around [sync.Pool](https://golang.org/pkg/sync/#Pool) that is used
to store [\*bytes.Buffer.](https://golang.org/pkg/bytes/#Buffer)

## Usage

```go

// Create a new pool
p := bufpool.New()

// Get bytes.Buffer from pool
buf := p.Get()

buf.WriteString("Hello, World!")
fmt.Println(buf.String())

// Put the buffer back into the pool
p.Put(buf)

```

*I've implemented this enough times that this small dependency comes in 
handy in reducing boilerplate code.*
