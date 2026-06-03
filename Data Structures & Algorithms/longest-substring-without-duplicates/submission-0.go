func lengthOfLongestSubstring(s string) int {

	l:=0
	maxlength:=0

	m:=make(map[byte]int)

	for r:=0;r<len(s);r++{

		if index,found:=m[s[r]];found && index>=l{
			l=index+1
		}

		m[s[r]]=r

		windowlength:=(r-l)+1
		if windowlength>maxlength{
			maxlength=windowlength
		}

	}

	return maxlength

}
