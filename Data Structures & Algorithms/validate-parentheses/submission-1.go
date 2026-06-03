func isValid(s string) bool {
	stack:=[]rune{}

	for _,val:=range s{
		switch val{
			case '(','{','[':
			stack=append(stack,val)

			case ')','}',']':
			if len(stack)==0{
				return false
			}
			top:=stack[len(stack)-1]

			if (val==')' && top !='(')||(val=='}' && top !='{')||(val==']' && top !='['){
				return false
			}

			stack=stack[:len(stack)-1]
			   
		}
	}

	return len(stack) == 0
    
}
