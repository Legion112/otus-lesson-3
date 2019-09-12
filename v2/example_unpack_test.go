package unpack_test

import (
	"fmt"

	unpack "github.com/mirecl/otus-lesson-3/v2"
)

// This example demonstrates error result function
func ExampleUnpack_withErrors() {
	if data, err := unpack.Unpack("df09"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}
	// Output:
	// invalid string
}

// This example demonstrates good result
func ExampleUnpack_withNoErrors() {
	if data, err := unpack.Unpack("df4"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}
	// Output:
	// dffff
}
