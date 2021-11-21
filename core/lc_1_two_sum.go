/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

*/

package core

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		key := target - nums[i]
		if _, ok := m[key]; ok {
			return []int{m[key], i}
		}
		m[nums[i]] = i
	}
	return nil

}
