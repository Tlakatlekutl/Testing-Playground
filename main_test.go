package main

import (
	"testing"
	"os"
	"io/ioutil"
)

type TestPair struct {
	Input int
	Output string
}

var tests = []TestPair{
	{18, "10: 1\n5: 1\n2: 1\n1: 1\n"},
	{9, "5: 1\n2: 2\n"},
}

func TestCalc_change(t *testing.T) {
	for _, test := range tests {
		stdOld := os.Stdout
		r, stdTo, _ := os.Pipe()
		os.Stdout = stdTo

		calc_change(test.Input) //call func here

		stdTo.Close()
		result, _ := ioutil.ReadAll(r)
		os.Stdout = stdOld

		answer := test.Output
		if string(result) != answer {
			t.Errorf("Function`s output '%s' dont match correct answer '%s'", result, answer)
		}
	}

}
