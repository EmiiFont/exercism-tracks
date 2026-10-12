package sublist

// Relation type is defined in relations.go file.
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
		list1 := l1[index:]

		if ContainsList(list1, l2) {
			return RelationSuperlist
		}
	}
	for index := range l2 {
		list2 := l2[index:]
		if ContainsList(list2, l1) {
			return RelationSublist
		}
	}
	return RelationUnequal
}
