package avltree

// return negative, 0, postive on a < b, a == b, a > b
type Comparator func(a, b interface{}) int

type Node struct {
	Key      interface{}
	Value    interface{}
	Parent   *Node
	Children [2]*Node
	b        int8
}

type Tree struct {
	Root       *Node
	Comparator Comparator
	size       int
}

func rotate(c int8, s *Node) *Node {
	a := (c + 1) / 2
	r := s.Children[a]
	s.Children[a] = r.Children[a^1]
	if s.Children[a] != nil {
		s.Children[a].Parent = s
	}
}

func singlerot(c int8, s *Node) *Node {
	s.b = 0
}

func doublerot(c int8, s *Node) *Node {

}

func putFix(c int8, t **Node) bool {
	s := *t
	if s.b == 0 {
		s.b = c
		return true
	}

	if s.b == -c {
		s.b = 0
		return false
	}

	if s.Children[(c+1)/2].b == c {

	} else {

	}
	*t = s
	return false
}

/*
key: key
value: value
p: parent node
qp: pointer to root node
*/
func (t *Tree) put(key, value interface{}, p *Node, qp **Node) bool {
	q := *qp
	if q == nil {
		t.size++
		*qp = &Node{Key: key, Value: value, Parent: p}
		return true
	}

	c := t.Comparator(key, q.Key)
	if c == 0 {
		q.Key = key
		q.Value = value
		return false
	}

	if c < 0 {
		c = -1
	} else {
		c = 1
	}
	a := (c + 1) / 2
	fix := t.put(key, value, q, &q.Children[a])
	if fix {
		return putFix()
	}
	return false
}

func (t *Tree) Put(key, value interface{}) {
	t.put(key, value, nil, &t.Root)
}
