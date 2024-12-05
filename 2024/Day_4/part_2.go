package day4_24

import (
	"os"
	"strconv"

	shared "github.com/VladanT3/Advent_of_Code"
)

func Part2() {
	data := GetData()
	checker := 0
	xmas_counter := 0

	for i, arr := range data {
		for j := range arr {
			if arr[j] == 'A' && (i > 0 && i < len(data)-1) && (j > 0 && j < len(arr)-1) {
				//looking for \
				if data[i-1][j-1] == 'M' && data[i+1][j+1] == 'S' {
					checker++
				} else if data[i+1][j+1] == 'M' && data[i-1][j-1] == 'S' {
					checker++
				}

				//looking for /
				if data[i-1][j+1] == 'M' && data[i+1][j-1] == 'S' {
					checker++
				} else if data[i+1][j-1] == 'M' && data[i-1][j+1] == 'S' {
					checker++
				}

				if checker == 2 {
					xmas_counter++
				}
				checker = 0
			}
		}
	}

	out, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("\nPart 2: " + strconv.Itoa(xmas_counter))
	out.Sync()
}
