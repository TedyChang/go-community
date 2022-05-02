package util

func HashPassword(pw *string) []byte {
	return []byte(*pw)
}

func MatchPassword(pw string, pw2 string) bool {
	return pw == pw2
}
