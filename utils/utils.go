package utils

//FindSlices procurar palavra no slice
func FindSlices(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}
