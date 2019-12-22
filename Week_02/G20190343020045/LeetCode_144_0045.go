func preorderTraversal(root *TreeNode) []int {

	var sum []int
	echo(root,&sum)
	return sum
}

func echo(root *TreeNode,sum *[]int)  {
	if root == nil {
		return
	}
	*sum = append(*sum, root.Val)
	echo(root.Left,sum)
	echo(root.Right,sum)
}
