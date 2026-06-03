type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for _,val:=range strs{
       sb.WriteString(strconv.Itoa(len(val)))
		sb.WriteString("%")
		sb.WriteString(val)
	}
return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
result := []string{}
	i := 0

	for i < len(encoded) {
		j := i

		
		for encoded[j] != '%' {
			j++
		}

		length, _ := strconv.Atoi(encoded[i:j])

		word := encoded[j+1 : j+1+length]
		result = append(result, word)

		i = j + 1 + length
	}

	return result
}
