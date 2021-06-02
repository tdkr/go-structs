package avltree

// return negative, 0, postive on a < b, a == b, a > b
type Comparator func(a, b interface{}) int

type Node struct {
	Key      interface{}
	Value    interface{}
	Parent   *Node
	Children [2]*Node
	b        int8 //balance
}

type Tree struct {
	Root       *Node
	Comparator Comparator
	size       int
}

/*
rotate s from left if c == -1
rotate s from right if c == 1
*/
func rotate(c int8, s *Node) *Node {
	a := (c + 1) / 2
	r := s.Children[a]
	s.Children[a] = r.Children[a^1]
	if s.Children[a] != nil {
		s.Children[a].Parent = s
	}
	r.Children[a^1] = s
	r.Parent = s.Parent
	s.Parent = r
	return r
}

func singlerot(c int8, s *Node) *Node {
	s.b = 0
	s = rotate(c, s)
	s.b = 0
	return s
}

func doublerot(c int8, s *Node) *Node {
	a := (c + 1) / 2
	r := s.Children[a]
	s.Children[a] = rotate(-c, s.Children[a])
	p := rotate(c, s)

	switch {
	case p.b == c:
		s.b = -c
		r.b = 0
	case p.b == -c:
		s.b = 0
		r.b = c
	default:
		s.b = 0
		r.b = 0
	}

	p.b = 0
	return p
}

/*
c: -1 if put to left child, 1 if put to right child
t: current node
return true on balance needed fix to parent node
*/
func putFix(c int8, t **Node) bool {
	s := *t
	if s.b == 0 {
		s.b = c
		return true
	}

	// insert to another child node of current node, balance stay the same for parent node
	if s.b == -c {
		s.b = 0
		return false
	}

	if s.Children[(c+1)/2].b == c {
		/*
			left left -> rotate right
			right right -> rotate left
		*/
		s = singlerot(c, s)
	} else {
		/*
			left right -> rotate left then right
			right left -> rotate right then left
		*/
		s = doublerot(c, s) // left right or right left
	}
	*t = s
	return false
}

/*
key: key
value: value
p: parent node
qp: current node
return false when node with same key already exists in tree
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
	a := (c + 1) / 2 // put to left child or right child
	fix := t.put(key, value, q, &q.Children[a])
	if fix {
		return putFix(int8(c), qp) // update balance
	}
	return false
}

func (t *Tree) Put(key, value interface{}) {
	t.put(key, value, nil, &t.Root)
}
