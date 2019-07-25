package iteration

// Repeat a letter n times
func Repeat(letter string, n int) (repeat string) {
	for i := 0; i < n; i++ {
		repeat += letter
	}
	return
}
