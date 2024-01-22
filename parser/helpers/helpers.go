package helpers

func AreAllInRange(values []int, min int, max int) bool {
	for _, value := range values {
		if value < min || value > max {
			return false
		}
	}
	return true
}
