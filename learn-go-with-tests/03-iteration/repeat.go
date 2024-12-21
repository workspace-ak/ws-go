package iteration

// Repeat 's' string five times
func Repeat(s string) string {
	repeated := ""
	for i := 0; i < 5; i++ {
		repeated += s
	}
	return repeated
}
