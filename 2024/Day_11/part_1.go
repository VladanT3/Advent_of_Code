package day11_24

import (
	"io"
	"math"
	"os"
	"strconv"

	shared "github.com/VladanT3/Advent_of_Code"
)

type Node struct {
	Value int
	Next  *Node
}

func Part1() {
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

	stone_head := &Node{
		Value: stones[0],
	}

	curr_stone := stone_head

	for i := 1; i < len(stones); i++ {
		curr_stone.Next = &Node{
			Value: stones[i],
		}

		curr_stone = curr_stone.Next
	}

	stone_count := len(stones)
	for range 25 {
		blink(stone_head, &stone_count)
	}

	out, err := os.Create("output.txt")
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("Part 1: " + strconv.Itoa(stone_count))
	shared.ErrCheck(err)
	out.Sync()
}

func blink(curr *Node, stone_count *int) {
	if curr.Value == 0 {
		curr.Value = 1
		if curr.Next != nil {
			blink(curr.Next, stone_count)
		}
		return
	}

	num := curr.Value
	digit_counter := 0
	for num > 0 {
		num /= 10
		digit_counter++
	}

	if digit_counter%2 == 0 {
		second_half := &Node{
			Value: 0,
		}
		second_half_digits := 0
		limit := digit_counter / 2
		for digit_counter > limit {
			second_half.Value += int(math.Pow(10, float64(second_half_digits))) * (curr.Value % 10)
			curr.Value /= 10
			second_half_digits++
			digit_counter--
		}

		second_half.Next = curr.Next
		curr.Next = second_half
		*stone_count++

		if curr.Next.Next != nil {
			blink(curr.Next.Next, stone_count)
		}
		return
	}

	curr.Value *= 2024
	if curr.Next != nil {
		blink(curr.Next, stone_count)
	}
	return
}
