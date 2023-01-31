package test

// AssessArrays compares two arrays and assess if they are equal
func AssessArrays[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}
