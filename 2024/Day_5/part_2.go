package day5_24

import (
	"os"
	"strconv"

	shared "github.com/VladanT3/Advent_of_Code"
)

func Part2() {
	rules := GetRules()
	data := GetData()
	correct := true
	sum := 0
	incorrect_data := [][]int{}

	for _, arr := range data {
		for j := range arr {
			after := rules[arr[j]]
			for k := j + 1; k < len(arr); k++ {
				_, ok := after[arr[k]]
				if !ok {
					correct = false
					incorrect_data = append(incorrect_data, arr)
					break
				}
			}
			if !correct {
				break
			}
		}
		correct = true
	}

	for _, arr := range incorrect_data {
		local_rules := make(map[int]map[int]bool)
		for j := range arr {
			if local_rules[arr[j]] == nil {
				local_rules[arr[j]] = make(map[int]bool)
			}
			after := rules[arr[j]]
			for k := range arr {
				if arr[k] != arr[j] {
					_, ok := after[arr[k]]
					if ok {
						local_rules[arr[j]][arr[k]] = true
					}
				}
			}
		}

		new_arr := []int{}
		ordering_map := make(map[int]int)
		for j := range arr {
			after, ok := local_rules[arr[j]]
			if !ok {
				ordering_map[0] = arr[j]
			} else {
				ordering_map[len(after)] = arr[j]
			}
		}

		for j := 0; j < len(arr); j++ {
			new_arr = append(new_arr, ordering_map[j])
		}

		sum += new_arr[(len(new_arr)-1)/2]
	}

	out, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("\nPart 2: " + strconv.Itoa(sum))
	shared.ErrCheck(err)
	out.Sync()
}
