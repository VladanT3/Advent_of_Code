package day11_24

import (
	"fmt"
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

	stone_count := len(stones)
	blinks := 0
	i := 0
	for blinks < 75 {
		for i < len(stones) {
			if stones[i] == 0 {
				stones[i] = 1
				i++
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

				if i < len(stones)-1 {
					in := -1
					out := second_half
					stones = append(stones, stones[len(stones)-1])
					for j := i + 1; j < len(stones)-1; j++ {
						in = out
						out = stones[j]
						stones[j] = in
					}
				} else {
					stones = append(stones, second_half)
				}
				stone_count++
				i += 2
				break
			}

			stones[i] *= 2024
			i++
		}
		if i >= len(stones) {
			blinks++
			i = 0
		}
	}

	fmt.Println(stone_count)

	out, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("\nPart 2: " + strconv.Itoa(stone_count))
	shared.ErrCheck(err)
	out.Sync()
}
