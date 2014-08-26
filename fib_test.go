package fib

import (
	"strconv"
	"strings"
	"testing"
)

func TestFib(t *testing.T) {
	var a []string

	for v := range Fib(10) {
		a = append(a, strconv.Itoa(v))
	}

	actual := strings.Join(a, ", ")
	if actual != "0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55" {
		t.Errorf("error: %v", actual)
	}
}
