func longestConsecutive(nums []int) int {

	set:=make(map[int]bool)

	for _,val:=range nums{
		set[val]=true
	}

	max:=0
	for num:=range set{

		if !set[num-1]{
			length:=1
            current:=num
		for set[current+1]{
			length++
			current++
		}
        if length>max{
			max=length
		}

		}
		

	}

	return max

}
