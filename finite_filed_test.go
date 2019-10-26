package galmug

import "testing"

func TestFieldElem(t *testing.T) {
	cases := []struct {
		desc string
		n    int
		p    int
		want int
	}{
		{"element 7 in 13 finite field is 7", 7, 13, 7},
	}

	for _, test := range cases {
		t.Run(test.desc, func(t *testing.T) {
			got, error := NewFieldElem(test.n, test.p)
			if error != nil {
				t.Errorf("Error occured with the value n=%q, p=%q", test.n, test.p)
			}
			if got.n != test.want {
				t.Errorf("got %q, want %q", got, test.want)
			}
			if got.p != test.p {
				t.Errorf("got %q, want %q", got, test.p)
			}
		})
	}
}

func TestFieldElemBoundary(t *testing.T) {
	cases := []struct {
		desc string
		n    int
		p    int
	}{
		{"Element 15 in F13 is invalid", 15, 13},
		{"Element 13 in F13 is invalid", 13, 13},
		{"Element -1 in F13 is invalid", -1, 13},
	}

	for _, test := range cases {
		t.Run(test.desc, func(t *testing.T) {
			got, error := NewFieldElem(test.n, test.p)
			if got != nil {
				t.Errorf("Boundary was not checked for the value n=%q, p=%q", test.n, test.p)
			}

			if error == nil {
				t.Errorf("No proper error was provided")
			}
		})
	}

}
