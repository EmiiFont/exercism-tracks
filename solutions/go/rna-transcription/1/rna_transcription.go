package strand

func ToRNA(dna string) string {
	dnaToRnaMap := map[string]string{
		"G": "C",
		"C": "G",
		"T": "A",
		"A": "U",
	}

	res := ""
	for _, val := range dna {
		val, ok := dnaToRnaMap[string(val)]
		if !ok {
			return ""
		}
		res += val
	}

	return res
}
