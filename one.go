package one

//one does what one wants
func one(b int, a int) bool {
	if b != 2 {
		if b+a > 10 {
			return true
		}
		return false
	}
	return true
}

func two(b int) bool {
	if b > 1 {
		return true
	}
	return false
}
