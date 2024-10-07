package plot

import (
	"fmt"
	"github.com/itsabgr/ge"
	_ "github.com/itsabgr/ge"
	"github.com/xlab/treeprint"
	_ "github.com/xlab/treeprint"
)

type treeNode struct {
	Tree    treeprint.Tree
	Count   int
	IsMulti bool
}

func sprintErr(err error) string {
	if s, ok := err.(fmt.Stringer); ok {
		return s.String()
	}
	return err.Error()
}

func Tree(root error) interface {
	Bytes() []byte
	fmt.Stringer
} {

	if root == nil {
		return nil
	}

	treeMap := map[int]*treeNode{}

	for depth, err := range ge.Walk(root) {
		node := &treeNode{}
		switch err.(type) {
		case ge.UnwrapError:
			if depth == 0 {
				node.Tree = treeprint.NewWithRoot(err)
			} else {
				node.Tree = treeMap[depth-1].Tree.AddBranch(err)
			}
		case ge.UnwrapErrors:
			node.IsMulti = true
			if depth == 0 {
				node.Tree = treeprint.NewWithRoot(0)
			} else {
				node.Tree = treeMap[depth-1].Tree.AddBranch(0)
			}

		default:
			if depth == 0 {
				node.Tree = treeprint.NewWithRoot(sprintErr(err))
			} else {
				node.Tree = treeMap[depth-1].Tree.AddNode(sprintErr(err))
			}
		}
		treeMap[depth] = node
		if prev := treeMap[depth-1]; prev != nil && prev.IsMulti {
			prev.Count += 1
			treeMap[depth-1].Tree.SetValue(prev.Count)
		}
	}
	return treeMap[0].Tree
}
