### CStruct 

This small package helps to copy values from a source struct to another struct with similar fields
and types. This is helpful for when there is need to copy a struct you have no control over. 

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

Set `cstruct.StrictMode(false|true)` to enable or disable case sensitive during field comparision.