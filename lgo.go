// Package to make strings in some commons
// console info formats using ANSI colors
package lgo

import "fmt"

// Prints data using a logger
func Log(logger logger, data any) {
	fmt.Print(logger(data))
}
