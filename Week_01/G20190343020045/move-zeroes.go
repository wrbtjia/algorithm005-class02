func moveZeroes(nums []int)  {
    k:=0
	for i:=0;i< len(nums);i++  {
		if nums[i] != 0 {
			if i != k {
				nums[k],nums[i] = nums[i],nums[k]
				k++
			}else {
				k++
			}
		}

	}
}