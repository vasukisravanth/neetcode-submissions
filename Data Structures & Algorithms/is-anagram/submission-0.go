func isAnagram(s string, t string) bool {

	if len(s)!=len(t){
		return false
	}
m:=make(map[rune]int)
n:=make(map[rune]int)
for _,val:=range s{
	m[val]++
}
for _,val2:=range t{
	n[val2]++
}

for i,_:=range m{
	if m[i]!=n[i]{
		return false
	}
}

return true



}
