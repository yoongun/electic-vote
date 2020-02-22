package lacan

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

func TestFieldElemCompare(t *testing.T) {
	elem1_7, _:= NewFieldElem(1, 7)
	elem2_7, _:= NewFieldElem(2, 7)
	elem1_11, _:= NewFieldElem(1, 11)
	elem2_13, _:= NewFieldElem(2, 13)
	cases := []struct {
		desc 	string
		elem1	FieldElem
		elem2 	FieldElem
		want	bool
	}{
		{"elem1 == elem2 (true case)", *elem1_7, *elem1_7, true},
		{"elem1 == elem2 (false case)", *elem2_7, *elem1_7, false},
		{"elem1 == elem2 (false case)", *elem1_7, *elem1_11, false},
		{"elem1 == elem2 (false case)", *elem1_7, *elem2_7, false},
		{"elem1 == elem2 (true case)", *elem2_13, *elem2_13, true},
	}

	for _, test := range cases {
		t.Run(test.desc, func(t *testing.T) {
			got1 := test.elem1.Eq(test.elem2)
			got2 := test.elem1.Neq(test.elem2)

			if got1 != test.want || got2 == test.want {
				t.Errorf("Comparison does not match the value desired")
			}
		})
	}
}
