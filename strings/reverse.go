package strings

func Reverse(sl string) string {
	runes := []rune(sl)
	reversed := []rune{}
	for i := len(runes) - 1; i >= 0; i-- {
		reversed = append(reversed, runes[i])
	}
	return string(reversed)
}
