func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
	sort.Ints(s)
	child := 0
	cook := 0
	for child <len(g) && cook < len(s)  {
		if g[child] <= s[cook] {
			child ++
		}
		cook ++
	}
	return child
}