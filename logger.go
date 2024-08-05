package lgo

import "fmt"

// Prototype of logger function
type logger func(any) string

// Returns input string in
// debug format (Green)
func Debug(data any) string{
	return fmt.Sprintf(makeFormatString(DEBUG_TAG), data)
}

// Returns input string in
// error format (Red)
func Error(data any) string{
	return fmt.Sprintf(makeFormatString(ERROR_TAG), data)
}

// Returns input string in
// info format (Teal)
func Info(data any) string{
	return fmt.Sprintf(makeFormatString(INFO_TAG), data)
}