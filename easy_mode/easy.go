package easy_mode

import "sort"

// 两数之和，只有一个解
func TwoSum(nums []int, target int) []int {
	lth := len(nums)
	if lth <= 1 {
		return []int{}
	}
	mp := make(map[int]int)
	for i := 0; i < lth; i++ {
		left := target - nums[i]
		v, ok := mp[left]
		if ok {
			return []int{v, i}
		}
		mp[nums[i]] = i
	}
	return []int{}
}

// 两数之和，只有一个解
// 第二种解法

type TwoSumItem struct {
	Value int
	Index int
}
func TwoSum2(nums []int, target int) []int {
	lth := len(nums)
	if lth <= 1 {
		return []int{}
	}
	newItems := make([]TwoSumItem, lth)
	for i := 0; i < lth; i++ {
		item := TwoSumItem{
			Value: nums[i],
			Index: i,
		}
		newItems[i] = item
	}
	// 先排序
	sort.Slice(newItems, func(i, j int) bool {
		return newItems[i].Value < newItems[j].Value
	})
	for i,j := 0, lth-1; i < j; {
		sum := newItems[i].Value + newItems[j].Value
		if sum == target {
			return []int{newItems[i].Index, newItems[j].Index}
		} else if sum < target {
			i++
		} else {
			j--
		}
	}
	return []int{}
}

// 3数之和，多个解
func ThreeSum(nums []int, target int) [][]int {
	result := [][]int{}
	lth := len(nums)
	if lth <= 2 {
		return result
	}
	// 排序
	//sort.Ints(nums)
	for i := 0; i <= lth - 3; i++ {
		if i == 0 || nums[i] != nums[i - 1] {
			j := i + 1
			k := lth - 1
			for j < k {
				sum := nums[i] + nums[j] + nums[k]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[k]})
					j++
					for j < k && nums[j] == nums[j-1] {
						j++
					}
					k--
					for j < k && nums[k] == nums[k+1] {
						k--
					}
				} else if sum < target {
					j++
				} else {
					k--
				}
			}
		}
	}
	return result
}

// 4数之和，多个解
func FourSum(nums []int, target int) [][]int {
	result := [][]int{}
	lth := len(nums)
	if lth <= 3 {
		return result
	}
	// 先排序
	sort.Ints(nums)
	for i := 0; i <= lth - 4; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			left := ThreeSum(nums[i+1:], target - nums[i])
			for _, v := range left {
				item := []int{nums[i]}
				item = append(item, v...)
				result = append(result, item)
			}
		}
	}
	return result
}

// 组合数
// candidate为不重复的正整数，每个candidate可使用多次
// dp，自下而上，从dp(1)计算到dp(target)
func CombinationSum(candidates []int, target int) [][]int {
	dp := make([][][]int, target+1)
	for i := 0; i <= target; i++ {
		dp[i] = make([][]int,0)
	}
	// 先排序
	sort.Ints(candidates)
	for i := 1; i <= target; i++ {
		for _, candidate := range candidates {
			if candidate > i {
				break
			} else if candidate == i {
				dp[i] = append(dp[i], []int{candidate})
				break
			}

			for _, list := range dp[i - candidate] {
				if candidate <= list[0] {
					tmp := []int{candidate}
					tmp = append(tmp, list...)
					dp[i] = append(dp[i], tmp)
				}
			}
		}
	}
	return dp[target]
}

// 组合数
// candidate为不重复的正整数，每个candidate可使用多次
// dp，每个candidate可以构成dp(candidate)~dp(target)的解
func CombinationSum2(candidates []int, target int) [][]int {
	dp := make([][][]int, target+1)
	for i := 0; i <= target; i++ {
		dp[i] = make([][]int,0)
	}
	// 先排序
	sort.Ints(candidates)
	for _, candidate := range candidates {
		for i := candidate; i <= target; i++ {
			if i == candidate {
				dp[i] = append(dp[i], []int{candidate})
				continue
			}
			for _, list := range dp[i - candidate] {
				tmp := []int{candidate}
				tmp = append(tmp, list...)
				dp[i] = append(dp[i], tmp)
			}
		}
	}
	return dp[target]
}
