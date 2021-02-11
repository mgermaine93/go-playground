// First test writing attempt in Go...

package addition

import "testing"

// Test functions always need to 1) Be Capitalized, 2) Be passed "t *testing.T"
func TestAndOneAndOne(t *testing.T) {

	actual := addOneAndOne()

	if actual != 2 {
		t.Errorf("Expected 2, but got %v", actual)
	}

}
