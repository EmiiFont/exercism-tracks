package protein

import (
	"errors"
	"fmt"
)

type ErrInvalidBase struct {
	Err error
}

func (e *ErrInvalidBase) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Err)
}

func FromRNA(rna string) ([]string, error) {
	panic("Please implement the FromRNA function")
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

	if _, ok := codonMap[codon]; !ok {
		return "", &ErrInvalidBase{Err: errors.New("sss")}
	}
	return codonMap[codon], nil
}
