package piscine

func StrRev(s string) string {
	char := []rune(s)

	for i, j := 0, len(char)-1; i < len(char)/2; i, j = i+1, j-1 {
		temp := char[j]
		char[j] = char[i]
		char[i] = temp
	}
	return string(char)
}
