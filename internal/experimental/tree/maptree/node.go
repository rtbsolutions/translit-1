package maptree

import "github.com/mxmCherry/translit/internal/experimental/tree"

// Node represents tree node with map lookup.
type Node struct {
	Val      []byte
	Children map[byte]*Node
}

// Value implements tree.Node.
func (n *Node) Value() []byte {
	return n.Val
}

// LookupChild implements tree.Node.
func (n *Node) LookupChild(b byte) tree.Node {
	return n.Children[b]
}

// HasChildren implements tree.Node.
func (n *Node) HasChildren() bool {
	return len(n.Children) > 0
}

// SetValue implements tree.BuilderNode.
func (n *Node) SetValue(v []byte) {
	n.Val = v
}

// EnsureChild implements tree.BuilderNode.
func (n *Node) EnsureChild(b byte) tree.BuilderNode {
	c := n.Children[b]
	if c == nil {
		c = &Node{}
		if n.Children == nil {
			n.Children = map[byte]*Node{
				b: c,
			}
		} else {
			n.Children[b] = c
		}
	}
	return c
}
