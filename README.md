# Thread-safe circular queue for Go
[![GoDoc](https://godoc.org/github.com/Mottl/circularqueue?status.svg)](https://godoc.org/github.com/Mottl/circularqueue)


## Installation
```sh
go get github.com/Mottl/circularqueue
```

## Example
```go
package main
  
import (
    "fmt"
    "github.com/Mottl/circularqueue"
)

func main() {
    queue := circularqueue.NewQueue(64)
    queue.Push(100)
    queue.Push(101)
    queue.Push(102)
    queue.Push(103)
    queue.PopAt(1)     // 101, nil
    queue.Pop()        // 103, nil
    fmt.Println(queue) // Queue (len=2, cap=64) [ 100, 102 ]
}
```

## License
Use of this package is governed by MIT license
that can be found in the LICENSE file.
