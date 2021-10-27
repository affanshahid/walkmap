# Walkmap

A Go library to iterate over potentially nested map keys using the visitor pattern

## Installing

```sh
go get -u github.com/affanshahid/walkmap
```

## Usage

```go
package main

import (
	"reflect"

	"github.com/affanshahid/walkmap"
)

func main() {
	data := map[interface{}]interface{}{
		"1": 1,
		"2": 2,
		"3": 3,
		"nested": map[interface{}]interface{}{
			"4": 4,
			"5": 5,
			"6": 6,
			"evenMoreNested": map[interface{}]interface{}{
				"7": 7,
				"8": 8,
				"9": 9,
			},
		},
	}

	walkmap.Walk(data, func(keyPath []interface{}, value interface{}, kind reflect.Kind) {
		// keyPath is the slice of keys used to arrive at the current node
		// value is the current node's value
		// kind describes the nature of the value, gotten using `reflect.TypeOf(value).Kind()`
		//...
	})
}
```
