package day13_24

import (
	"fmt"
	"io"
	"math/big"
	"os"

	shared "github.com/VladanT3/Advent_of_Code"
)

func Part2() {
	data := GetDataPart2()
	a := big.NewInt(0)
	b := big.NewInt(0)
	tokens := big.NewInt(0)

	for _, machine := range data {
		x1 := big.NewInt(machine[0])
		y1 := big.NewInt(machine[2])
		x2 := big.NewInt(machine[1])
		y2 := big.NewInt(machine[3])
		result1 := new(big.Int).Add(big.NewInt(machine[4]), big.NewInt(10000000000000))
		result2 := new(big.Int).Add(big.NewInt(machine[5]), big.NewInt(10000000000000))
		b = new(big.Int).Div(
			new(big.Int).Sub(new(big.Int).Mul(result1, x2), new(big.Int).Mul(result2, x1)),
			new(big.Int).Sub(new(big.Int).Mul(y1, x2), new(big.Int).Mul(y2, x1)))
		//b = ((machine[4] * machine[1]) - (machine[5] * machine[0])) / ((machine[2] * machine[1]) - (machine[3] * machine[0]))
		a = new(big.Int).Div(
			new(big.Int).Sub(result1, new(big.Int).Mul(b, y1)),
			x1)
		//a = (machine[4] - b*machine[2]) / machine[0]

		func1 := new(big.Int).Add(new(big.Int).Mul(x1, a), new(big.Int).Mul(y1, b))
		func2 := new(big.Int).Add(new(big.Int).Mul(x2, a), new(big.Int).Mul(y2, b))
		//if machine[0]*a+machine[2]*b == machine[4] && machine[1]*a+machine[3]*b == machine[5] {
		if func1.Cmp(result1) == 0 && func2.Cmp(result2) == 0 {
			tokens = new(big.Int).Add(tokens,
				new(big.Int).Add(
					new(big.Int).Mul(a, big.NewInt(3)), b))
			//tokens += a*3 + b
		}
		//}
	}

	fmt.Println(tokens)

	out, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("\nPart 2: " + tokens.String())
	shared.ErrCheck(err)
	out.Sync()
}

func GetDataPart2() [][]int64 {
	f, err := os.Open("input.txt")
	shared.ErrCheck(err)
	defer f.Close()

	eof, err := f.Seek(0, io.SeekEnd)
	shared.ErrCheck(err)
	cursor, err := f.Seek(0, io.SeekStart)
	shared.ErrCheck(err)

	char := make([]byte, 1)
	out := [][]int64{}
	arr := []int64{}
	var num int64 = 0

	for cursor < eof {
		_, err = f.Read(char)
		shared.ErrCheck(err)
		cursor++

		if int(char[0]) >= 48 && int(char[0]) <= 57 {
			num = (num * 10) + (int64(char[0]) - 48)
		} else if num != 0 {
			arr = append(arr, num)
			num = 0
		}

		if len(arr) == 6 {
			out = append(out, arr)
			arr = []int64{}
		}
	}

	return out
}
