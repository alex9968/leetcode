package test

import (
	"reflect"
	"testing"
)

func TestA(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expect:%v, got: %v", want, got)
	}
}

func TestGroup(t *testing.T) {
	type testType struct {
		input string
		sep   string
		want  []string
	}

	tests := map[string]testType{
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong seq":   {input: "a:b:c", sep: ",", want: []string{"a", "b", "c"}},
		"more seq":    {input: "abcd", sep: "bc", want: []string{"a", "b", "c"}},
		"leading seq": {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expect:%v, got: %v", tc.want, got)
			}
		})
	}
}
