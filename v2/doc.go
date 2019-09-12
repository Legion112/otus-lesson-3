/*
Package unpack function performs primitive unpacking of a string containing repeated characters, for example:
	*"a4bc2d5e" => "aaaabccddddde"
	*"abcd" => "abcd"
	*"45" = > "" (invalid string)

Optional: support for the escape sequences
	*qwe\4\5 => qwe45 (*)
	*qwe\45 => qwe44444 (*)
	*qwe\\5 => qwe\\\\\ (*)

Example used:
	package main

	import (
		"fmt"
		unpack "github.com/mirecl/otus-lesson-3/v2"
	)

	func main() {
		if data, err := pack.Unpack("a4"); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(data)
		}
	}
*/
package unpack
