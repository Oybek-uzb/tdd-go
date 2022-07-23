package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpextedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{
				"Chris",
				"New York",
			},
			[]string{"Chris", "New York"},
		},
		{
			"Struct with non-string field",
			struct {
				Name string
				Age  int
			}{
				"Chris",
				89,
			},
			[]string{"Chris"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpextedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpextedCalls)
			}
		})
	}
}
