package day7_24

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	shared "github.com/VladanT3/Advent_of_Code"
)

func Part2() {
	f, err := os.ReadFile("input.txt")
	shared.ErrCheck(err)

	data := strings.Split(string(f), "\n")
	sum := 0

	for i := 0; i < len(data)-1; i++ {
		row := strings.Split(data[i], ": ")
		target, err := strconv.Atoi(row[0])
		shared.ErrCheck(err)

		nums_str := strings.Split(row[1], " ")
		nums := []int{}
		for j := range nums_str {
			num, err := strconv.Atoi(nums_str[j])
			shared.ErrCheck(err)
			nums = append(nums, num)
		}

		if IsValidPart2(nums, target, 0, 0) {
			sum += target
		}
	}

	fmt.Println(sum)

	out, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("\nPart 2: " + strconv.Itoa(sum))
	shared.ErrCheck(err)
	out.Sync()
}

func IsValidPart2(nums []int, target int, idx int, curr_sum int) bool {
	if idx == len(nums)-1 {
		return curr_sum+nums[idx] == target || curr_sum*nums[idx] == target || ConcatNums(curr_sum, nums[idx]) == target
	}

	return IsValidPart2(nums, target, idx+1, curr_sum+nums[idx]) || IsValidPart2(nums, target, idx+1, curr_sum*nums[idx]) || IsValidPart2(nums, target, idx+1, ConcatNums(curr_sum, nums[idx]))
}

func ConcatNums(num1 int, num2 int) int {
	concat_num, err := strconv.Atoi(fmt.Sprintf("%d%d", num1, num2))
	shared.ErrCheck(err)

	return concat_num
}
