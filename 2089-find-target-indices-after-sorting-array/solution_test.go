package findtargetindicesaftersortingarray

import (
	"testing"
)

func TestSearchInt(t *testing.T) {
	idxx := targetIndices([]int{1, 2, 5, 2, 3}, 2)
	t.Log(idxx)
}
