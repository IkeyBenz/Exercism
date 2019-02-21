package protein

// FromCodon does stuff
func FromCodon(codon string) (string, error) {
	var protein string

	switch codon {
	case "AUG":
		protein = "Methionine"
	case "UUU", "UUC":
		protein = "Phenylalanine"
	case "UUA", "UUG":
		protein = "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		protein = "Serine"
	case "UAU", "UAC":
		protein = "Tyrosine"
	case "UGU", "UGC":
		protein = "Cysteine"
	case "UGG":
		protein = "Tryptophan"
	case "UAA", "UAG", "UGA":
		protein = ""
	}

	return protein, nil
}

// FromRNA takes a string of RNA characters and returns the Codons it contains
func FromRNA(s string) ([]string, error) {

	var proteins []string
	temp := ""

	for i, char := range s {

		temp += string(char)

		if i+1%3 == 0 {
			protein, err := FromCodon(temp)
			if err != nil {
				panic(err)
			}
			if protein == "STOP" {
				return proteins, nil
			}
			proteins = append(proteins, protein)
			temp = ""
		}

	}

	return proteins, nil

}
