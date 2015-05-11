# Skip List in Golang #

[![Build Status](https://travis-ci.org/huandu/skiplist.svg?branch=master)](https://travis-ci.org/huandu/skiplist)
[![GoDoc](https://godoc.org/github.com/huandu/skiplist?status.svg)](https://godoc.org/github.com/huandu/skiplist)

Skip list is a kind of ordered map and can store any value inside. See [skip list](http://en.wikipedia.org/wiki/Skip_list) wikipedia page to learn more about this data structure.

Highlights in this implementation:

* Support custom compare function so that any type can be used as key.
* Key sort order can be changed quite easily.
* Rand source and max level can be changed per list. It can be useful in performance critical scenarios.

## How To Use ##

Install this package through `go get`.

    go get github.com/huandu/skiplist

Use it as following.

```go
package main

import (
    "fmt"
    "github.com/huandu/skiplist"
)

func main() {
    // Create a skip list with int key.
    list := skiplist.New(skiplist.Int)

    // Add some values. Value can be anything.
    list.Set(12, "hello world")
    list.Set(34, 56)

    // Get element by index.
    elem := list.Get(34) // Value is stored in elem.Value.
    fmt.Println(elem.Value)
    next := elem.Next()  // Get next element.
    fmt.Println(next.Value)

    // Or get value directly just like a map
    val, ok := list.GetValue(34)
    fmt.Println(val, ok)

    // Remove an element by index.
    list.Remove(34)
}
```

## License ##

This library is licensed under MIT license. See LICENSE for details.
