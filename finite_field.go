package lacan

import "errors"

type FieldElem struct {
	n int
	p int
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
