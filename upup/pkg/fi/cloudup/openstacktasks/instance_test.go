package openstacktasks

import "testing"

func Test_ParseQuantityString(t *testing.T) {
	grid := []struct {
		i      string
		expect int
		pass   bool
	}{
		{i: "15", expect: 15, pass: true},
		{i: "34g", expect: 34, pass: true},
		{i: "5t", expect: 5000, pass: true},
		{i: ".1t", expect: 100, pass: true},
		{i: ".13t", expect: 130, pass: true},
		{i: ".1g", expect: 0, pass: true},
		{i: ".1k", expect: 0, pass: false},
		{i: "1s", expect: 0, pass: false},
		{i: "i4", expect: 0, pass: false},
	}

	for _, test := range grid {
		val, err := parseQuantityString(test.i)
		if test.pass {
			if err != nil {
				t.Errorf("error thrown, but expected to pass: %v, %v", test.i, err)
			}

			if test.expect != val {
				t.Errorf("no error thrown but results don't match, expected: %v, actual: %v", test.expect, val)
			}
		}

		if !test.pass && err == nil {
			t.Errorf("expected an error but function passed %v", err)
		}
	}
}
