package ge

import (
	"fmt"
	"github.com/itsabgr/ge/plot"
	"testing"
)

func TestTree(t *testing.T) {
	var errs []error
	for i := range 10 {
		errs = append(errs, fmt.Errorf("err%d", i))
	}
	err := Wrap(Join(Join(Wrap(Wrap(Join(errs[2], errs[1], errs[8]), errs[3]), errs[4]), errs[5]), errs[6]), errs[7])
	tree := plot.Tree(err)
	fmt.Print(tree.String())
}
