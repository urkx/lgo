package lgo

import (
	"fmt"
	"testing"
)

func TestLogUInt(t *testing.T) {
	Log(Debug, 9)
}

func TestLogUInt8(t *testing.T) {
	var x uint = 8
	Log(Debug, x)
	Log(Info, x)
	Log(Error, x)
}

func TestLogUInt16(t *testing.T) {
	var x uint16 = 16
	Log(Debug, x)
	Log(Info, x)
	Log(Error, x)
}

func TestLogUInt32(t *testing.T) {
	var x uint32 = 32
	Log(Debug, x)
	Log(Info, x)
	Log(Error, x)
}

func TestLogUInt64(t *testing.T) {
	var x uint64 = 64
	Log(Debug, x)
	Log(Info, x)
	Log(Error, x)
}

func TestLogInt(t *testing.T) {
	Log(Debug, 9)
}

func TestLogInt8(t *testing.T) {
	var x int = 8
	Log(Debug, x)
	Log(Info, x)
	Log(Error, x)
}

func TestLogInt16(t *testing.T) {
	var x int16 = 16
	Log(Debug, x)
	Log(Info, x)
	Log(Error, x)
}

func TestLogInt32(t *testing.T) {
	var x int32 = 32
	Log(Debug, x)
	Log(Info, x)
	Log(Error, x)
}

func TestLogInt64(t *testing.T) {
	var x int64 = 64
	Log(Debug, x)
	Log(Info, x)
	Log(Error, x)
}

func TestLogString(t *testing.T) {
	Log(Debug, "String log")
}

func TestLogStringFormatted(t *testing.T) {
	Log(Debug, fmt.Sprintf("Formatted string: %d", 5))
}

func TestLogBasicArray(t *testing.T) {
	x := [5]int{1, 2, 3, 4, 5}
	Log(Info, x)
}

func TestLogStruct(t *testing.T) {
	type X struct {
		foo int
		bar int
	}
	x := X{foo: 1, bar: 2}
	Log(Debug, x)
}

func TestLogStructArray(t *testing.T) {
	type X struct {
		foo int
		bar int
	}

	x := [2]X{{foo: 1, bar: 2}, {foo: 3, bar: 4}}
	Log(Error, x)
}