package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid - 10/2", 10.0, 2.0, 5.0, false},
	{"invalid - 10/0", 10.0, 0.0, 0.0, true},
	{"expect-fraction", -1.0, -777.0, 0.0012870013, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := divide(tt.dividend, tt.divisor)
			if tt.isErr {
				if err == nil {
					t.Error("Did not get an error when we should have")
				}
			} else {
				if err != nil {
					t.Error("Got an error when we should not have", err.Error())
				}
			}
			if result != tt.expected {
				t.Errorf("Got %f, expected %f", result, tt.expected)
			}
		})
	}
}
