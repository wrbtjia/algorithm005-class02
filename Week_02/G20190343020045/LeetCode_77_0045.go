func combine(n int, k int) [][]int {
    var result [][]int
	var temp []int
	combinations(n,k,1,temp,&result)
	return result
}

func combinations(n int, k int,start int,temp []int,result *[][]int)  {
	if k == len(temp){
		t := make([]int, len(temp))
		copy(t, temp)
		*result = append(*result, t)
	}
	for i := start;i <= n;i++{
		temp = append(temp, i)
		combinations(n,k,i+1,temp,result)
		temp = temp[:len(temp)-1]
	}
}