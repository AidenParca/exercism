package dna

import "fmt"

type Histogram map[rune]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, ncl := range d {
		_, ok := h[ncl]
		if ok {
			h[ncl]++
		} else {
			return nil, fmt.Errorf("invalid nucleotide '%c' found", ncl)
		}
	}
	return h, nil
}