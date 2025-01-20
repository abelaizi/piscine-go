package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		if root.Left != nil {
			root.Left.Parent = root
		}
	} else if data == root.Data {
	} else {
		root.Right = BTreeInsertData(root.Right, data)
		if root.Right != nil {
			root.Right.Parent = root
		}
	}
	return root
}
