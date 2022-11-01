package main

import (
	"testing"
)

func Test_SayHello_ValidArgument(t *testing.T) {
	inputs := []struct {
		name   string
		result string
	}{
		{name: "Golang", result: "Hello Golang"},
		{name: "C#", result: "Hello C#"},
		{name: "Python", result: "Hello Python"},
	}

	for _, item := range inputs {

		result := sayHello(item.name)
		if result != item.result {
			t.Errorf("\"sayHello('%s')\" failed, expected -> %v, got -> %v", item.name, item.result, result)
		} else {
			t.Logf("\"sayHello('%s')\" succeded, expected -> %v, got -> %v", item.name, item.result, result)
		}
	}
}

func BenchmarkSayHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sayHello("Golang")
	}
}
