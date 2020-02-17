package array_test

import (
	"fmt"

	"github.com/frozentech/array"
)

func ExampleInArray() {
	sourceString := []string{
		"A",
		"B",
	}

	fmt.Println(array.InArray("A", sourceString))
	fmt.Println(array.InArray("B", sourceString))
	fmt.Println(array.InArray("Z", sourceString))

	sourceInt := []int{
		1,
		2,
	}

	fmt.Println(array.InArray(1, sourceInt))
	fmt.Println(array.InArray(2, sourceInt))
	fmt.Println(array.InArray(0, sourceInt))

	sourceFloat := []float64{
		1.2,
		2.4,
	}

	fmt.Println(array.InArray(1.2, sourceFloat))
	fmt.Println(array.InArray(2.4, sourceFloat))
	fmt.Println(array.InArray(2.41, sourceFloat))

	// Output:
	// true 0
	// true 1
	// false -1
	// true 0
	// true 1
	// false -1
	// true 0
	// true 1
	// false -1
}

func ExampleKeyExist() {
	source := map[string]string{
		"a": "A",
		"b": "B",
		"c": "C",
	}

	fmt.Println(array.KeyExist("a", source))
	fmt.Println(array.KeyExist("b", source))
	fmt.Println(array.KeyExist("c", source))
	fmt.Println(array.KeyExist("z", source))
	fmt.Println(array.KeyExist("A", source))
	fmt.Println(array.KeyExist("B", source))
	fmt.Println(array.KeyExist("C", source))

	// Output:
	// true
	// true
	// true
	// false
	// false
	// false
	// false
}
