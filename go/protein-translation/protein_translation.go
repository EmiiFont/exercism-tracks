package protein

import (
	"errors"
)

var (
	ErrInvalidBase = errors.New("invalid base")
	ErrStop        = errors.New("ErroStop")
)

func FromRNA(rna string) ([]string, error) {
	res := []string{}
	for i := 0; i < len(rna); i += 3 {
		codon := rna[i : i+3]
		// UGG-AGA-AUU-AAU-GGU-UU
		val, err := FromCodon(codon)
		if err != nil {
			if err == ErrInvalidBase {
				return nil, err
			}
			break
		}
		res = append(res, val)
	}
	return res, nil
}

func FromCodon(codon string) (string, error) {
	codonMap := map[string]string{
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

	value, ok := codonMap[codon]
	if !ok {
		return "", ErrInvalidBase
	}

	if value == "STOP" {
		return "", ErrStop
	}

	return codonMap[codon], nil
}
