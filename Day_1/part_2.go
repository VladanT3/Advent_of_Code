package day1

import (
	"os"
	"strconv"
)

func Part2() {
	arr1, arr2 := GetLists()
	freq_map := make(map[int]int)
	freq_sum := 0

	for i := range arr2 {
		_, ok := freq_map[arr2[i]]
		if !ok {
			freq_map[arr2[i]] = 1
		} else {
			freq_map[arr2[i]]++
		}
	}

	for i := range arr1 {
		freq, ok := freq_map[arr1[i]]
		if ok {
			freq_sum += arr1[i] * freq
		}
	}

	f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	ErrCheck(err)
	defer f.Close()

	_, err = f.WriteString("\nPart 2: " + strconv.Itoa(freq_sum))
	f.Sync()
}
