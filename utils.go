package utils

func Contains(a []string, x string) bool {
	for _, s := range a {
		if s == x {
			return true
		}
	}
	return false
}
func ContainsInt(a []int, x int) bool {
	for _, s := range a {
		if s == x {
			return true
		}
	}
	return false
}
