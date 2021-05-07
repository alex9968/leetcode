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

func TestB(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expect:%v, got: %v", want, got)
	}
}
