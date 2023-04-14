package tree

import (
	"sort"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/ivanilves/travelgrunt/pkg/directory/tree/node"
)

// Tree is a hierarchical representation of the directory contents
type Tree struct {
	nodes map[string]node.Node

	levels []map[string]string
}

func getLevels(paths []string) (levels []map[string]string) {
	for _, path := range paths {
		elements := strings.Split(path, "/")

		previous := ""
		current := ""

		for idx, el := range elements {
			if previous != "" {
				current = previous + "/" + el
			} else {
				current = el
			}

			if len(levels) < idx+1 {
				levels = append(levels, make(map[string]string, 0))
			}

			levels[idx][current] = el

			previous = current
		}
	}

	return levels
}

func getNodeParent(path string, parentLevel map[string]string) string {
	for parentPath := range parentLevel {
		if strings.HasPrefix(path, parentPath+"/") {
			return parentPath
		}
	}

	return ""
}

func getNodeChildren(path string, levels []map[string]string, idx int) (children []string) {
	if len(levels) <= idx+1 {
		return []string{}
	}

	childLevel := levels[idx+1]

	for childPath := range childLevel {
		if strings.HasPrefix(childPath, path+"/") {
			children = append(children, childPath)
		}
	}

	return children
}

func getNodes(levels []map[string]string, paths []string) (nodes map[string]node.Node) {
	nodes = make(map[string]node.Node, 0)

	var prevLevel map[string]string

	for idx, level := range levels {
		for path, name := range level {
			nodes[path] = node.NewNode(
				name,
				path,
				getNodeParent(path, prevLevel),
				getNodeChildren(path, levels, idx),
				idx,
				slices.Contains(paths, path),
			)
		}

		prevLevel = level
	}

	return nodes
}

func sortedKeys(items map[string]string) (keys []string) {
	keys = make([]string, len(items))

	c := 0
	for key := range items {
		keys[c] = key

		c++
	}

	sort.Strings(keys)

	return keys
}

// NewTree creates a new Tree struct from the passed list of directory paths
func NewTree(paths []string) Tree {
	levels := getLevels(paths)

	nodes := getNodes(levels, paths)

	return Tree{nodes: nodes, levels: levels}
}

// LevelCount gives us the number of hierarchical levels inside the tree
func (t Tree) LevelCount() int {
	return len(t.levels)
}

func (t Tree) levelItems(idx int) (items map[string]string) {
	if len(t.levels) <= idx+1 {
		return nil
	}

	items = make(map[string]string, len(t.levels[idx]))

	for path, name := range t.levels[idx] {
		items[name] = path
	}

	return items
}

// GetNode returns a node record for the given path
func (t Tree) GetNode(path string) node.Node {
	return t.nodes[path]
}

// LevelChildItems gives us a list of child items for the [parent] node located on the given level ID and on the given path
func (t Tree) LevelChildItems(idx int, parentPath string) (items map[string]string) {
	if len(t.levels) < idx+1 {
		return nil
	}

	if len(t.levels) < idx+2 {
		return map[string]string{}
	}

	if idx == -1 {
		return t.levelItems(0)
	}

	items = make(map[string]string, len(t.levels[idx+1]))

	for path, name := range t.levels[idx+1] {
		currentNode := t.GetNode(path)
		parentNode := t.GetNode(parentPath)

		if currentNode.IsAChildOf(parentNode) {
			items[name] = path
		}

		if parentNode.IsTerminal() && parentNode.HasChildren() {
			items["."] = parentPath
		}
	}

	return items
}

// LevelChildNames gives us a list of child names for the [parent] node located on the given level ID and on the given path
func (t Tree) LevelChildNames(idx int, parentPath string) []string {
	return sortedKeys(t.LevelChildItems(idx, parentPath))
}
