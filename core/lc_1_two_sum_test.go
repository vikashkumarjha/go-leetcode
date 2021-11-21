package core

import (
	"testing"
)

func TestTwoSum(t *testing.T) {

	t.Logf("running the two sum test")
	nums1 := []int{1, 2, 3, 4, 5}
	target1 := 9
	r := TwoSum(nums1, target1)
	e := []int{3, 4}
	if r[0] != e[0] || r[1] != e[1] {
		t.Errorf("got(%v) expected(%v)", r, e)
	}

	nums2 := []int{}
	target2 := 10

	r2 := TwoSum(nums2, target2)

	if r2 != nil {
		t.Errorf("got(%v) expected(%v)", r2, nil)
	}
}
