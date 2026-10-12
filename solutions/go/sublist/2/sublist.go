package sublist

// Relation type is defined in relations.go file.
func IsSubsequence(primary, secondary []int) bool {
	if len(primary) == 0 {
		return true
	}
	if len(secondary) == 0 {
		return false
	}
	if len(primary) > len(secondary) {
		return false
	}
	for i := 0; i < len(secondary); i++ {
		if primary[0] == secondary[i] {
			if len(primary) == 1 {
				return true
			}
			if len(primary) > 1 {
				return IsSubsequence(primary[1:], secondary[i+1:])
			}
		}
	}
	return false
}

func ContainsList(l1, l2 []int) bool {
	if len(l1) < len(l2) {
		return false
	}

	for index, val := range l2 {
		if l1[index] != val {
			return false
		}
	}
	return true
}

func AreEqual(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}

	for index, val := range l1 {
		if l2[index] != val {
			return false
		}
	}
	return true
}

func Sublist(l1, l2 []int) Relation {
	if AreEqual(l1, l2) {
		return RelationEqual
	}

	for index := range l1 {
		l1 = l1[index:]

		if ContainsList(l1, l2) {
			return RelationSuperlist
		}
	}
	for index := range l2 {
		l2 = l2[index:]
		if ContainsList(l2, l1) {
			return RelationSublist
		}
	}
	return RelationUnequal
}
