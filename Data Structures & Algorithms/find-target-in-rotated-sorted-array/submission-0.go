func search(nums []int, target int) int {

	m:=make(map[int]int)
	for i,val:=range nums{
       m[val]=i
	}

	if v, found := m[target]; found {
	return v
	}
	return -1

}
