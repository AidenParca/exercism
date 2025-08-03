package strand

import "strings"

func ToRNA(dna string) string {
	var rna strings.Builder
	rna.Grow(len(dna))
	for _, nucleotide := range dna {
		switch nucleotide {
		case 'G':
			rna.WriteRune('C')
		case 'C':
			rna.WriteRune('G')
		case 'T':
			rna.WriteRune('A')
		case 'A':
			rna.WriteRune('U')
		}
	}
	return rna.String()
}