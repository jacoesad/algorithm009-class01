package Week_03

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}

	if n == 1 {
		return root
	}

	rootInInorder := -1
	for index, number := range inorder {
		if number == root.Val {
			rootInInorder = index
			break
		}
	}

	if rootInInorder == -1 {
		return nil
	}

	leftPreorder := preorder[1 : 1+rootInInorder]
	rightPreorder := preorder[1+rootInInorder:]
	leftInorder := inorder[:rootInInorder]
	rightInorder := inorder[rootInInorder+1:]

	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)

	return root
}
