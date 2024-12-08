package day7_24

import (
	"fmt"
	"io"
	"math/big"
	"os"

	shared "github.com/VladanT3/Advent_of_Code"
)

func Part1() {
	data := GetData()

	for target, nums := range data {

	}
}

func GetData() map[*big.Int][]*big.Int {
	f, err := os.Open("input.txt")
	shared.ErrCheck(err)
	defer f.Close()

	eof, err := f.Seek(0, io.SeekEnd)
	shared.ErrCheck(err)
	i, err := f.Seek(0, io.SeekStart)
	shared.ErrCheck(err)

	char := make([]byte, 1)
	target := big.NewInt(0)
	set_target := false
	num := big.NewInt(0)
	nums := []*big.Int{}
	data := make(map[*big.Int][]*big.Int)

	for i < eof {
		_, err = f.Read(char)
		shared.ErrCheck(err)
		i++
		if string(char[0]) == "\n" {
			nums = append(nums, num)
			num = big.NewInt(0)
			data[target] = nums
			target = big.NewInt(0)
			nums = []*big.Int{}
			set_target = false
			continue
		}
		if char[0] == ':' {
			set_target = true
			_, err = f.Seek(1, io.SeekCurrent)
			shared.ErrCheck(err)
			i++
			continue
		}
		if !set_target {
			digit := big.NewInt(int64(char[0]) - 48)
			target = new(big.Int).Mul(target, big.NewInt(10))
			target = new(big.Int).Add(target, digit)
		} else {
			if char[0] != ' ' {
				digit := big.NewInt(int64(char[0]) - 48)
				num = new(big.Int).Mul(num, big.NewInt(10))
				num = new(big.Int).Add(num, digit)
			} else {
				nums = append(nums, num)
				num = big.NewInt(0)
			}
		}
	}

	return data
}
