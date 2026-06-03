func productExceptSelf(nums []int) []int {

	left:=make([]int,len(nums))
	right:=make([]int,len(nums))
		result:=make([]int,len(nums))

left[0]=1
	for i:=1;i<len(nums);i++{
    left[i]=left[i-1]*nums[i-1]
	}

	right[len(nums)-1]=1
	for j:=len(nums)-2;j>-1;j--{
		right[j]=right[j+1]*nums[j+1]
	}

	for k:=0;k<len(nums);k++{
		result[k]=left[k]*right[k]
	}

	return result


}
