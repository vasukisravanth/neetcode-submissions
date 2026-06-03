func isPalindrome(s string) bool {
   var result strings.Builder

	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {   // keeps only letters
			result.WriteRune(unicode.ToLower(ch))
		}
	}

	b:=result.String()

	l:=0
	r:=len(b)-1
	for l<r{
		if b[l]!=b[r]{
			return false
		}
		l++
		r--
	}

	return true
}
