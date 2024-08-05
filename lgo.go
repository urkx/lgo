// Package to make strings in some commons
// console info formats using ANSI colors
package lgo

import "fmt"

const(
	DEBUG_FORMAT string = "\033[1;32m[DEBUG]\033[0m: %s\n" // Green
	ERROR_FORMAT string = "\033[1;31m[ERROR]\033[0m: %s\n" // Red
	INFO_FORMAT string	= "\033[1;36m[INFO]\033[0m: %s\n" // Teal
)

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
