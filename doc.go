/*
Package unpack function performs primitive unpacking of a string containing repeated characters, for example:
	*"a4bc2d5e" => "aaaabccddddde"
	*"abcd" => "abcd"
	*"45" = > "" (invalid string)

Optional: support for the escape sequences
	*qwe\4\5 => qwe45 (*)
	*qwe\45 => qwe44444 (*)
	*qwe\\5 => qwe\\\\\ (*)
*/
package unpack
