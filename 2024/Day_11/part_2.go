package day11_24

import (
	"io"
	"math"
	"os"
	"strconv"

	shared "github.com/VladanT3/Advent_of_Code"
)

func Part2() {
	f, err := os.Open("input.txt")
	shared.ErrCheck(err)
	defer f.Close()

	eof, err := f.Seek(0, io.SeekEnd)
	shared.ErrCheck(err)
	cursor, err := f.Seek(0, io.SeekStart)
	shared.ErrCheck(err)

	char := make([]byte, 1)
	stones := []int{}
	num := 0

	for cursor < eof {
		_, err = f.Read(char)
		shared.ErrCheck(err)
		cursor++

		if char[0] == ' ' {
			stones = append(stones, num)
			num = 0
			continue
		}

		if int(char[0]) >= 48 && int(char[0]) <= 57 {
			num = (num * 10) + (int(char[0]) - 48)
		}
	}
	stones = append(stones, num)

	for range 75 {
		for i := range stones {
			if stones[i] == 0 {
				stones[i] = 1
				continue
			}

			num := stones[i]
			digit_counter := 0
			for num > 0 {
				num /= 10
				digit_counter++
			}

			if digit_counter%2 == 0 {
				second_half := 0
				second_half_digits := 0
				limit := digit_counter / 2
				for digit_counter > limit {
					second_half += int(math.Pow(10, float64(second_half_digits))) * (stones[i] % 10)
					stones[i] /= 10
					second_half_digits++
					digit_counter--
				}

				stones = append(stones, second_half)
				continue
			}

			stones[i] *= 2024
		}
	}

	out, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("\nPart 2: " + strconv.Itoa(len(stones)))
	shared.ErrCheck(err)
	out.Sync()
}
