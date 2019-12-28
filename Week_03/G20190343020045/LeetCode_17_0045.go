var kv = map[int]string {
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

func letterCombinations(digits string) []string {
	var res []string
	if len(digits) == 0 {
		return res
	}
	letterCombinationsCore(digits, 0, &res, "")
	return res
}

func letterCombinationsCore(digits string, index int, pRes *[]string, str string) {
	if index == len(digits) -1 {
		for _,v := range kv[int(digits[index] -'0')] {
			*pRes = append(*pRes, str+string(v))
		}
	}else {
		for _,v := range kv[int(digits[index]-'0')] {
			letterCombinationsCore(digits,index+1,pRes,str+string(v))
		}
	}


}