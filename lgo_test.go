package lgo

import (
	"fmt"
	"testing"
)

func TestDebug(t *testing.T) {
	fmt.Print(Debug("Hola en debug"))
}

func TestError(t *testing.T) {
	fmt.Print(Error("Algo ha fallado :-/"))
}

func TestInfo(t *testing.T) {
	fmt.Print(Info("Info"))
}
