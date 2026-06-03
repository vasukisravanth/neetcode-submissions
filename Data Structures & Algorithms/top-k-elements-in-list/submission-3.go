func topKFrequent(nums []int, k int) []int {

	m:=make(map[int]int)
	n:=make([]int,0,len(m))



	for _,val:=range nums{
		if v,exists:=m[val];exists{
			m[val]=v+1
			
		}else{
			m[val]=1
		}
	}

	for i:=range m{
		n=append(n,i)
	}

	 sort.Slice(n, func(i, j int) bool {
    return m[n[i]] > m[n[j]]
})


	return n[:k]



}
