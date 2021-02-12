// // First test writing attempt in Go...

package main

import "testing"
func TestAddOneAndOne(t *testing.T) {

	result := addOneAndOne()
	if result != 2 {
		t.Error("Expected", 2, "Got", result)
	}
}