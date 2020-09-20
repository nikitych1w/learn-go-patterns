package generator_test

import (
	"fmt"
	"github.com/nikitych1w/learn-go-patterns/Behavioral/generator"
	"testing"
)

func TestFib(t *testing.T) {
	for i := range generator.Fibonacci(1000) {
		fmt.Println(i)
	}
}
