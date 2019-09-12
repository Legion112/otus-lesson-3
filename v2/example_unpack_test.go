package unpack_test

import (
	"fmt"

	unpack "github.com/mirecl/otus-lesson-3/v2"
)

func ExampleUnpack() {
	if data, err := unpack.Unpack("df09"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}
}
