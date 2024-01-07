package main

import (
	"go_blind_75/arrayhashing"
	"testing"
)

func TestContainsDuplicates(t *testing.T) {
	var actual = arrayhashing.ContainsDuplicates([]int8{1, 2, 3, 1})
	if actual == false {
		t.Errorf("Expected %v, but got %v", true, actual)
	}
}
