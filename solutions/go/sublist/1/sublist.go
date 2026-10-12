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

func Sublist(l1, l2 []int) Relation {
	if len(l1) == 0 && len(l2) == 0 {
		return RelationEqual
	}
	if len(l1) == 0 {
		return RelationSublist
	}
	if len(l2) == 0 {
		return RelationSuperlist
	}
	if len(l1) == len(l2) {
		if IsSubsequence(l1, l2) {
			return RelationEqual
		}
	}
	if len(l1) > len(l2) {
		if IsSubsequence(l2, l1) {
			return RelationSuperlist
		}
	}
	if len(l1) < len(l2) {
		if IsSubsequence(l1, l2) {
			return RelationSublist
		}
	}

	return RelationUnequal
}
