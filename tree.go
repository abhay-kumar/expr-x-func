package exprxfunc

type Tree struct {
	Left  *Tree
	Value string
	Right *Tree
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) AddValue(value string) {
	t.Value = value
}

func (t *Tree) AddLeft(lt *Tree) {
	t.Left = lt
}

func (t *Tree) AddRight(rt *Tree) {
	t.Right = rt
}
