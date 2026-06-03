func groupAnagrams(strs []string) [][]string {
	m:=make(map[string][]string)

	for _,val:=range strs{

		chars:=[]rune(val)
		 sort.Slice(chars, func(i, j int) bool {
            return chars[i] < chars[j]
        })

        key := string(chars)

      m[key]=append(m[key],val)

	}
	s:=make([][]string,0)

	for _,val:=range m{
		s=append(s,val)
	}

	return s

}
