package node

// Node is an element Tree gets composed from (a path enriched with hierarchical and connection metadata)
type Node struct {
	name     string
	path     string
	parent   string
	children []string
	levelID  int
	terminal bool
}

// NewNode creates a new Node instance
func NewNode(name, path, parent string, children []string, levelID int, terminal bool) Node {
	return Node{
		name:     name,
		path:     path,
		parent:   parent,
		children: children,
		levelID:  levelID,
		terminal: terminal,
	}
}

// IsAChildOf tells us if current node is a child of another node
func (n Node) IsAChildOf(pn Node) bool {
	return n.parent == pn.path
}

// HasChildren tells us if that node has child nodes (i.e. if node is a parent itself)
func (n Node) HasChildren() bool {
	return len(n.children) > 0
}

// IsTerminal tells us if that node has can be used as a tree exit point
func (n Node) IsTerminal() bool {
	return n.terminal
}
