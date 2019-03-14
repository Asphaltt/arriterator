# arriterator
A simple iterator for Go array/slice

## Usage

```go
package main

import (
  arr "github.com/Asphaltt/arriterator"
  "fmt"
)

func main() {
  i, e := arr.New([]int{111, 222, 333})
  if e != nil {
    panic(e)
  }
  
  for i.HasNext() {
    fmt.Println(i.Next())
  }
}
```
