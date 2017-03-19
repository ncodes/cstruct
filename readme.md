### CStruct 

CStruct copies a source struct to another struct that share similar fields
and field types.  

#### Installation
```
go get github.com/ncodes/cstructs
```

#### Example

```go
package main

import "github.com/ncodes/cstructs"

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
    cstructs.Copy(&ben, &pete)
}
```

Set `cstruct.StrictMode(false|true)` to enable or disable case sensitivity during field comparision.
