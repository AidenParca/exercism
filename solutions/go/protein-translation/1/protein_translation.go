package protein

import "errors"

var (
	ErrStop        = errors.New("stop codon encountered")
	ErrInvalidBase = errors.New("invalid codon")
)
var codonToAminoAcid = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}
func FromCodon(codon string) (string, error) {
	aminoAcid, exists := codonToAminoAcid[codon]
	if !exists {
		return "", ErrInvalidBase
	}
	if aminoAcid == "STOP" {
		return "", ErrStop
	}
	return aminoAcid, nil
}
func FromRNA(rna string) ([]string, error) {
	var protein []string
    // 3 step jump , i+3 <= len helps it squeez to the end as well as knowing if there is less than 3 words left which is a false input error.
	for i := 0; i+3 <= len(rna); i += 3 {
		codon := rna[i : i+3]
		aminoAcid, err := FromCodon(codon)
		if err != nil {
			if errors.Is(err, ErrStop) {
				return protein, nil
			}
			return nil, err 
		}
		protein = append(protein, aminoAcid)
	}
	return protein, nil
}