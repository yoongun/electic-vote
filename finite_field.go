package evote

import "errors"

type FieldElem struct {
	n int
	p int
}

func (e FieldElem) Eq(t FieldElem) bool {
	if e.n == t.n && e.p == t.p {
		return true
	}
	return false
}

func (e FieldElem) Neq(t FieldElem) bool {
	return !e.Eq(t)
}

func NewFieldElem(n int, p int) (*FieldElem, error) {
	if n >= p {
		return nil, errors.New("n must be lower than p ")
	}
	if n < 0 {
		return nil, errors.New("n must be a positive integer ")
	}

	return &FieldElem{n, p}, nil
}
