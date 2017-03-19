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

#### StrictMode(true|false) (Default: true)

Enable  or disable case sensitivity during field comparision. If turned off, fields are normalized to 
lowercase before comparing with another.

#### CopySlice(src interface{}, dest interface{})

Copies a slice of struct to another slice of struct. Both slice must have the same length.If src is empty, `nil` is returned.

#### MakeSliceOf(of interface{}, size int)

A convenience method for creating a slice of an initialized type. Useful when there is need to create a slice of struct to use as destination when calling `CopySlice`.