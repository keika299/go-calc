// Copyright (c) 2017 keika299
//
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

package calc_test

import (
	"github.com/keika299/go-calc"
	"testing"
)

func TestRunInt(t *testing.T) {

	testFrags := []struct {
		input         string
		expect        int
		expectIsError bool
	}{
		{"invalid", 0, true},
		{"1", 1, false},
		{"2.0", 2, false},
	}

	for _, f := range testFrags {
		result, err := calc.RunInt(f.input)

		if err != nil != f.expectIsError {
			if err != nil {
				t.Errorf("Run should not return error.\nInput: %s\nError Message: %s",
					f.input, err.Error())
			}

			if err == nil {
				t.Errorf("Run should return error.\nInput: %s", f.input)
			}
		}

		if err == nil {
			if result != f.expect {
				t.Errorf("Run return value is mismatch.\nInput: %s\nExpect: %d\nActual: %d",
					f.input, f.expect, result)
			}
		}
	}
}

func TestRun(t *testing.T) {

	testFrags := []struct {
		input         string
		expect        float64
		expectIsError bool
	}{
		{"", 0.0, true},
		{"invalid", 0.0, true},
		{"1", 1.0, false},
		{"2.0", 2.0, false},
		{"-3.0", -3.0, false},
		{"+4.0", 4.0, false},
		{"1.0+2.0", 3.0, false},
		{"1.0 + 2.0", 3.0, false},
		{" 1.0+2.0 ", 3.0, false},
		{"7.0-2.0", 5.0, false},
		{"2.0*4.0", 8.0, false},
		{"9.0/3.0", 3.0, false},
		{"1.0/0.0", 0.0, true},
		{"-3.0+2.0", -1.0, false},
		{"-3.0*2.0", -6.0, false},
		{"1.0+2.0+4.0", 7.0, false},
		{"1.0+8.0-3.0", 6.0, false},
		{"4.0*2.0+2.0", 10.0, false},
		{"4.0+2.0*2.0", 8.0, false},
		{"4.0/2.0+2.0", 4.0, false},
		{"4.0+2.0/2.0", 5.0, false},
		{"4.0*2.0*2.0", 16.0, false},
		{"4.0/2.0/2.0", 1.0, false},
		{"4.0*2.0+2.0*3.0", 14, false},
		{"4.0/2.0-2.0*3.0", -4.0, false},
	}

	for _, f := range testFrags {
		result, err := calc.Run(f.input)

		if err != nil != f.expectIsError {
			if err != nil {
				t.Errorf("Run should not return error.\nInput: %s\nError Message: %s",
					f.input, err.Error())
			}

			if err == nil {
				t.Errorf("Run should return error.\nInput: %s", f.input)
			}
		}

		if err == nil {
			if result != f.expect {
				t.Errorf("Run return value is mismatch.\nInput: %s\nExpect: %f\nActual: %f",
					f.input, f.expect, result)
			}
		}
	}
}

func TestConditionalExpression(t *testing.T) {
	testFrags := []struct {
		input         string
		expect        bool
		expectIsError bool
	}{
		{"", false, true},
		{"invalid", false, true},
		{"1.0", false, true},
		{"1<2", true, false},
		{"1>2", false, false},
		{"1=2", false, false},
		{"1<=1", true, false},
		{"1>=1", true, false},
		{"1 = 1", true, false},
		{"1 = 2", false, false},
		{"1.1=1.1", true, false},
		{"1.1=2.1", false, false},
		{"1+2=3", true, false},
		{"2+2=3", false, false},
		{"1+7=3+5", true, false},
		{"2+2=3+4", false, false},
		{"12=3*4", true, false},
	}

	for _, f := range testFrags {
		result, err := calc.ConditionalExpression(f.input)

		if err != nil != f.expectIsError {
			if err != nil {
				t.Errorf("Run should not return error.\nInput: %s\nError Message: %s",
					f.input, err.Error())
			}

			if err == nil {
				t.Errorf("Run should return error.\nInput: %s", f.input)
			}
		}

		if err == nil {
			if result != f.expect {
				t.Errorf("Run return value is mismatch.\nInput: %s\nExpect: %f\nActual: %f",
					f.input, f.expect, result)
			}
		}
	}
}
