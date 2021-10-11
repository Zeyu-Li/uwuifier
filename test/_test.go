package test

import (
	"convert"
	"testing"
)

func TestWorking(t *testing.T) {
	input := ">++++++++++++++++++++++++++++++++++++++++++++----.>..,,[[]]]]]?><<"

	output := convert.Convert(input)

	if (output != "👉👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👇👇👇👇🥺👉🥺🥺😳😳😒😒😡😡😡😡😡🥴👉👈👈") {
		t.Fatalf("Does not match")
	}
}