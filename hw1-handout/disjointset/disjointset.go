package disjointset

// DisjointSet is the interface for the disjoint-set (or union-find) data
// structure.
// Do not change the definition of this interface.
type DisjointSet interface {
	// UnionSet(s, t) merges (unions) the sets containing s and t,
	// and returns the representative of the resulting merged set.
	UnionSet(int, int) int
	// FindSet(s) returns representative of the class that s belongs to.
	FindSet(int) int
}

// TODO: implement a type that satisfies the DisjointSet interface.

// NewDisjointSet creates a struct of a type that satisfies the DisjointSet interface.
type Node struct { //parent holds where integer is pointing to
	parent int
}

var m map[int]Node

func NewDisjointSet() DisjointSet {
	//panic("TODO: implement this!")

	m = make(map[int]Node) //make new map to act as Disjointset
	m[0] = Node{0}         //initialize the 0 integer in set, the key acts as value of integer

	return m[0] // need to return node
}

func (x Node) UnionSet(a int, b int) int {
	aparent, bparent := x.FindSet(a), x.FindSet(b) // find integer they are pointing to

	if aparent == bparent { // if points to same then in same set
		return aparent
	}

	m[bparent] = Node{aparent} //make one node have the other as a parent
	return aparent

}

func (x Node) FindSet(a int) int {

	_, ok := m[a]

	if ok == false { //check if map for a exists, if not then make one
		m[a] = Node{a}
	}
	aparent := m[a].parent
	for aparent != m[aparent].parent { //check until you find value have itself as parent meaning it is the representative

		m[aparent], aparent = Node{m[m[aparent].parent].parent}, m[aparent].parent //update the lower rank's parent to its parent's parent and move to next parent
	}

	return aparent // return representative
}
