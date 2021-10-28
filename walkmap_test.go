package walkmap

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/PaesslerAG/jsonpath"
)

func ExampleWalk() {
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

	concatenatePath := func(paths []interface{}) string {
		strPaths := make([]string, 0, len(paths))

		for _, path := range paths {
			strPaths = append(strPaths, path.(string))
		}

		return strings.Join(strPaths, ".")
	}

	flattened := map[string]interface{}{}
	paths := []string{}

	Walk(data, func(keyPath []interface{}, value interface{}, kind reflect.Kind) {
		pathStr := concatenatePath(keyPath)
		paths = append(paths, pathStr)
		flattened[pathStr] = value
	})

	sort.Strings(paths)

	for _, p := range paths {
		fmt.Printf("Key: %s, Value: %d\n", p, flattened[p])
	}

	// Output:
	// Key: 1, Value: 1
	// Key: 2, Value: 2
	// Key: 3, Value: 3
	// Key: nested.4, Value: 4
	// Key: nested.5, Value: 5
	// Key: nested.6, Value: 6
	// Key: nested.evenMoreNested.7, Value: 7
	// Key: nested.evenMoreNested.8, Value: 8
	// Key: nested.evenMoreNested.9, Value: 9
}

func TestWalkSimple(t *testing.T) {
	data := map[interface{}]interface{}{
		"1": 1,
		"2": 2,
		"3": 3,
	}

	got := map[interface{}]interface{}{}

	Walk(data, func(keyPath []interface{}, value interface{}, kind reflect.Kind) {
		got[keyPath[0]] = value
	})

	if len(got) != len(data) {
		t.Fatalf("Expected outmap to be of size %d, got %d", len(data), len(got))
	}

	for key, val := range got {
		if data[key] != val {
			t.Fatalf("Expected %s, got %s", data[key], val)
		}
	}
}

func TestWalkNested(t *testing.T) {
	data := map[interface{}]interface{}{
		"foo": "1",
		"bar": "2",
		"baz": "3",
		"nested": map[interface{}]interface{}{
			"fooNested": "4",
			"barNested": "5",
			"bazNested": "6",
			"evenMoreNested": map[interface{}]interface{}{
				"fooEvenMoreNested": "7",
				"barEvenMoreNested": "8",
				"bazEvenMoreNested": "9",
			},
		},
	}

	got := map[string]interface{}{}

	concatenatePath := func(paths []interface{}) string {
		strPaths := make([]string, 0, len(paths))

		for _, path := range paths {
			strPaths = append(strPaths, path.(string))
		}

		return strings.Join(strPaths, ".")
	}

	Walk(data, func(keyPath []interface{}, value interface{}, kind reflect.Kind) {
		pathStr := concatenatePath(keyPath)
		got[pathStr] = value
	})

	if len(got) != 9 {
		t.Fatalf("Expected outmap to be of size %d, got %d", 9, len(got))
	}

	for key, val := range got {
		exp, err := jsonpath.Get(key, data)
		if err != nil {
			t.Fatalf("Expected err to be nil, got: %s", err)
		}

		if exp != val {
			t.Fatalf("Expected %s, got %s", exp, val)
		}
	}
}

func TestWalkNestedMixedTypes(t *testing.T) {
	data := map[interface{}]interface{}{
		"foo": "1",
		"bar": "2",
		"baz": "3",
		"nested": map[string]interface{}{
			"fooNested": "4",
			"barNested": "5",
			"bazNested": "6",
			"evenMoreNested": map[interface{}]interface{}{
				"fooEvenMoreNested": "7",
				"barEvenMoreNested": "8",
				"bazEvenMoreNested": "9",
			},
		},
	}

	got := map[string]interface{}{}

	concatenatePath := func(paths []interface{}) string {
		strPaths := make([]string, 0, len(paths))

		for _, path := range paths {
			strPaths = append(strPaths, path.(string))
		}

		return strings.Join(strPaths, ".")
	}

	Walk(data, func(keyPath []interface{}, value interface{}, kind reflect.Kind) {
		pathStr := concatenatePath(keyPath)
		got[pathStr] = value
	})

	if len(got) != 9 {
		t.Fatalf("Expected outmap to be of size %d, got %d", 9, len(got))
	}

	for key, val := range got {
		exp, err := jsonpath.Get(key, data)
		if err != nil {
			t.Fatalf("Expected err to be nil, got: %s", err)
		}

		if exp != val {
			t.Fatalf("Expected %s, got %s", exp, val)
		}
	}
}
