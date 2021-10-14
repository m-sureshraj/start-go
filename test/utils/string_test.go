package utils

import (
	"fmt"
	"testing"
)

// how to run? "go test <package path>"
// e.g. go test start-go/test/utils

// go test command looks for files that are ending with "_test.go"
// and run all the functions which are start with "Test..."

// go test <path> -v : will print detailed summary of the test results
// go help test - for more detail

// func TestEmptyElement(t *testing.T) {
// 	list := []string{}
// 	want := ""
// 	got := Join(list)

// 	if want != got {
// 		t.Errorf(errorString(list, got, want))
// 	}
// }

// func TestSingleElement(t *testing.T) {
// 	list := []string{"foo"}
// 	want := "foo"
// 	got := Join(list)

// 	if want != got {
// 		t.Errorf(errorString(list, got, want))
// 	}
// }

// func TestTwoElements(t *testing.T) {
// 	list := []string{"foo", "bar", "baz"}

// 	want := "foo, bar and baz"
// 	got := Join(list)

// 	if want != got {
// 		t.Errorf(errorString(list, got, want))
// 	}
// }

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("Join(%#v) = \"%s\", want \"%s\"", list, got, want)
}

type testData struct {
	input []string
	want  string
}

func TestJoin(t *testing.T) {
	tests := []testData{
		{input: []string{}, want: ""},
		{input: []string{"foo"}, want: "foo"},
		{input: []string{"foo", "bar"}, want: "foo and bar"},
		{input: []string{"foo", "bar", "baz"}, want: "foo, bar, and baz"},
	}

	for _, test := range tests {
		got := Join(test.input)
		if test.want != got {
			t.Errorf(errorString(test.input, got, test.want))
		}
	}
}
