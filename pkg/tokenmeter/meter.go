package tokenmeter

// Estimate returns a rough token count for text (stub).
func Estimate(text string) int {
	if text == "" {
		return 0
	}
	return len(text) / 4
}