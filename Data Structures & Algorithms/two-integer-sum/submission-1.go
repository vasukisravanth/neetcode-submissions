func twoSum(nums []int, target int) []int {

	m:=make(map[int]int)

	for j,val:=range nums{
		m[val]=j
	}

	for i,val:=range nums{
		k:=target-val
		if id,exists:=m[k];exists && i!=id{
		 return []int{i, id}
		}
	}

	return nil
    
}
