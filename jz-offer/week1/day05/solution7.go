package day05

func inorderSuccessor(p *TreeNode) *TreeNode {
	// 右子树非空,找到右子树中最左边的树
	if p.Right != nil{
		p = p.Right
		for p.Left != nil{
			p = p.Left
		}
		return p
	}else{
		/*
		res    -->    o
					 /
					o
		             \
				      o
				       \
		                p
		*/
		for p.Father != nil && p.Father.Right == p {
			p = p.Father
		}
		return p.Father
	}
}