### CStruct 

CStruct allow you to copy from a source struct to another struct with similar fields
and types. It makes use of the `reflect` package. 

#### Installation
```
go get github.com/ncodes/cstruct
```

#### Example

```go
package main

import "github.com/ncodes/cstruct"

// A
type A struct {
	Name  string
	Age   int
	About []byte
}

// B
type B struct {
	Name  string
	Age   int
	About []byte
}

func main(){ 
    
    ben := A{
		Name:  "benedith",
		Age:   20,
		About: []byte("Ben is a great sports person"),
	}

    var pete B
    Copy(&ben, &pete)
}
```

Set `cstruct.StrictMode(false|true)` to enable or disable case sensitivity during field comparision.
