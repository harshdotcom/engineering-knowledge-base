# Slices

*  An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

* A slice does not store any data, it just describes a section of an underlying array. Changing the elements of a slice modifies the corresponding elements of its underlying array. Other slices that share the same underlying array will see those changes.

### Slice Insertion

## Range (continued)

When using `range` in Go, you can skip either the index or the value by assigning it to `_` (blank identifier).

### Skip the value (keep the index)
```go
for i, _ := range pow {
    // use i (index)
}
```


## Maps

A map in Go stores key-value pairs. It allows you to quickly look up a value using a key.

### Declaration
```go
var m map[string]int
m = make(map[string]int)

m["apple"] = 5
m["banana"] = 3
```

## Map Literals

Map literals are a way to create and initialize a map with values in a single statement.  
They are similar to struct literals, but **keys are required** for each value.

### Syntax

```go
map[KeyType]ValueType{
    key1: value1,
    key2: value2,
}


//

package main

import "fmt"

func main() {
    scores := map[string]int{
        "Alice": 90,
        "Bob":   85,
        "Carol": 92,
    }

    fmt.Println(scores)
}

```



