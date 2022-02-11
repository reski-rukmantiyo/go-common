package tools

// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
// Only works with array
func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

// Contains tells whether a contains x.
// Only works with array
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
