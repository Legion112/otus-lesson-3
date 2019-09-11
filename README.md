# lesson-3 (Unpack)

[![image](https://img.shields.io/badge/godoc-reference-blue)](http://godoc.org/github.com/mirecl/otus-lesson-3)
![image](https://img.shields.io/badge/coverage-100%25-green)

### Задание
Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:

* "a4bc2d5e" => "aaaabccddddde"
* "abcd" => "abcd"
* "45" => "" (некорректная строка)

Дополнительное задание: поддержка escape - последовательности
* `qwe\4\5` => `qwe45` (*)
* `qwe\45` => `qwe44444` (*)
* `qwe\\5` => `qwe\\\\\` (*)

### Documentation
* [API Reference](http://godoc.org/github.com/mirecl/otus-lesson-3)

### Installation

    go get github.com/mirecl/otus-lesson-2

## Examples
```go
package main

import (
	"fmt"

	unpack "github.com/mirecl/otus-lesson-3"
)

func main() {
	fmt.Println(unpack.Parser("a4bc2d5e"))
	fmt.Println(unpack.Parser2("a4bc2d5e"))
}
```

## Benchmark

Распаковка (CPU i7-6800):
```Go
`a4b\4\5c0d5e\52\\3qwe\4\5`
```
Результат:
```bash
goos: windows
goarch: amd64
pkg: github.com/mirecl/otus-lesson-2
BenchmarkParser-12     	 1000000	      1673 ns/op	    1200 B/op	      32 allocs/op
BenchmarkParser2-12    	 1000000	      1263 ns/op	    1456 B/op	       8 allocs/op
```