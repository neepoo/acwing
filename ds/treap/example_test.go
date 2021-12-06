package treap

import "testing"

func TestTreap(t *testing.T) {
	tp := NewTreap()
	tp.insert(str("Cabbage"), 77)
	tp.insert(str("Beer"), 76)
	tp.insert(str("Eggs"), 129)
	tp.insert(str("Bacon"), 95)
	tp.insert(str("Butter"), 86)
}
