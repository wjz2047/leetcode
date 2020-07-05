package easy_mode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	input := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(TwoSum(input, target))
}

func TestTwoSum2(t *testing.T) {
	fmt.Println("this is TwoSum2")
	input := []int{3, 2, 4}
	target := 6
	fmt.Println(TwoSum2(input, target))
}

func TestThreeSum(t *testing.T) {
	input := []int{-1, 0, 1, 2, -1, -4}
	target := 0
	fmt.Println(ThreeSum(input, target))
}

func TestFourSum(t *testing.T) {
	input := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println(FourSum(input, target))
}

func TestCombinationSum1(t *testing.T) {
	fmt.Println("this is TestCombinationSum")
	candidates := []int{35,29,32,40,44,33,39,23,20,36,42,22,48,25,47,26,37,21,27,41,46,49,30,43,28,34,31,24,38}
	target := 72
	fmt.Println(CombinationSum(candidates, target))
}

func TestCombinationSum2(t *testing.T) {
	fmt.Println("this is TestCombinationSum2")
	candidates := []int{35,29,32,40,44,33,39,23,20,36,42,22,48,25,47,26,37,21,27,41,46,49,30,43,28,34,31,24,38}
	target := 72
	fmt.Println(CombinationSum2(candidates, target))
}

func TestCoinChange(t *testing.T) {
	coins := []int{1,2,5}
	amount := 11
	fmt.Println(CoinChange(coins, amount))
}

func TestCoinChange2(t *testing.T) {
	coins := []int{1,2,5}
	amount := 11
	fmt.Println(CoinChange2(coins, amount))
}
