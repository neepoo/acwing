package treap

type SortedPriorityQueue interface {
	Top() interface{}
	Peek() interface{}
	Insert(element interface{}, priority int)
	Remove(element interface{})
	Update(element interface{}, newPriority int)
	Contains(element interface{}) bool
	Min()
	Max()
}

type Comparable interface {
	Less(other interface{}) bool
	Equal(other interface{}) bool
	Bigger(other interface{}) bool
}

type str string

func (s str) Less(other interface{}) bool {
	return s < other.(str)
}

func (s str) Equal(other interface{}) bool {
	return s == other.(str)
}

func (s str) Bigger(other interface{}) bool {
	return s > other.(str)
}

type Node struct {
	key Comparable // 按照这个字段排序(bst)

	priority float64 // 按照这个字段维护堆的性质

	Left   *Node
	Right  *Node
	Parent *Node
}

func NewNode(key Comparable, priority float64) *Node {
	return &Node{key: key, priority: priority}
}

// 更新一个节点的左子树
func (node *Node) setLeft(left *Node) {
	node.Left = left
	if node != nil {
		//fmt.Println("22222")
		if left != nil {
			left.Parent = node

		}
	}

}

// 更新一个节点的右子树
func (node *Node) setRight(right *Node) {
	node.Right = right
	if node != nil {
		if right != nil {
			right.Parent = node

		}
	}
	return
}

func (node *Node) Search(targetKey interface{}) *Node {
	t := node
	if t == nil {
		return nil
	}
	if t.key.Equal(targetKey) {
		return t
	} else if t.key.Bigger(targetKey) {
		return t.Left.Search(targetKey)
	} else {
		return t.Right.Search(targetKey)
	}
}

type Treap struct {
	root *Node
}

func NewTreap() *Treap {
	return &Treap{}
}

// RightRotate 对于x节点进行右旋, x是其父节点的左节点(此时x的priority小于其父节点的,不满足小根堆的性质)
// 1.将x换到其父节点的位置
// 2.将x的右节点设置为y(x.parent)的的左节点
// 3. 将x的右节点设置为y
func (t *Treap) RightRotate(x *Node) {
	y := x.Parent // x的父亲节点
	if y.Left != x {
		panic("Right rotate must happen at left node")
	}
	p := y.Parent
	if p != nil {
		if p.Left == y {
			p.setLeft(x)
		} else {
			p.setRight(x)
		}
	} else {
		//说明之前p是根节点
		t.root = x
	}
	// 现在y节点是没有接在树上的
	// step1: 将x的右节点接到y的左节点
	// step1: 将x的右节点设置为y
	y.setLeft(x.Right)
	x.setRight(y)
}

// LeftRotate 左旋x
func (t *Treap) LeftRotate(x *Node) {
	if x == nil || x == t.root {
		panic("left rotate: x is nil or x is root")
	}
	y := x.Parent
	if y.Right != x {
		panic("left rotate")
	}
	p := y.Parent
	if p != nil {
		if p.Left == y {
			p.setLeft(x)
		} else {
			p.setRight(x)
		}
	} else {
		t.root = x
	}
	y.setRight(x.Left)
	x.setLeft(y)
}

func (t *Treap) insert(key Comparable, priority float64) {
	newNode := NewNode(key, priority)
	node := t.root
	var parent *Node = nil
	for node != nil {
		parent = node
		// 这里左边右边应该都是可以的
		if node.key.Less(key) || node.key.Equal(key) {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	if parent == nil {
		t.root = newNode
		return
	} else if key.Less(parent.key) || key.Equal(parent.key) {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}
	newNode.Parent = parent
	for newNode.Parent != nil && newNode.priority < newNode.Parent.priority {
		if newNode == newNode.Parent.Left {
			t.RightRotate(newNode)
		} else {
			t.LeftRotate(newNode)
		}
	}
	if newNode.Parent == nil {
		t.root = newNode
	}
}
