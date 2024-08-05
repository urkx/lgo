// Package to make strings in some commons
// console info formats using ANSI colors
package lgo

import "fmt"

type formatter func(string) string

// Returns input string in
// debug format (Green)
func Debug(data string) string{
	return fmt.Sprintf(DEBUG_FORMAT, data)
}

// Returns input string in
// error format (Red)
func Error(data string) string{
	return fmt.Sprintf(ERROR_FORMAT, data)
}

// Returns input string in
// info format (Teal)
func Info(data string) string{
	return fmt.Sprintf(INFO_FORMAT, data)
}

// Prints data using a format function
func Log(format formatter, data string) {
	fmt.Print(format(data))
}
