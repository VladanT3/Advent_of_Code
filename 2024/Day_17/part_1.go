package day17_24

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	shared "github.com/VladanT3/Advent_of_Code"
)

func Part1() {
	a, b, c, program := GetData()
	reg_a := big.NewInt(int64(a))
	reg_b := big.NewInt(int64(b))
	reg_c := big.NewInt(int64(c))
	i := 0
	output := []*big.Int{}

	for i < len(program) {
		if program[i] == 0 {
			if program[i+1] < 4 {
				reg_a = new(big.Int).Div(reg_a, new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(program[i+1])), nil))
				//a /= int(math.Pow(2, float64(program[i+1])))
			} else {
				if program[i+1] == 4 {
					reg_a = new(big.Int).Div(reg_a, new(big.Int).Exp(big.NewInt(2), reg_a, nil))
					//a /= int(math.Pow(2, float64(a)))
				} else if program[i+1] == 5 {
					reg_a = new(big.Int).Div(reg_a, new(big.Int).Exp(big.NewInt(2), reg_b, nil))
					//a /= int(math.Pow(2, float64(b)))
				} else if program[i+1] == 6 {
					reg_a = new(big.Int).Div(reg_a, new(big.Int).Exp(big.NewInt(2), reg_c, nil))
					//a /= int(math.Pow(2, float64(c)))
				}
			}
		} else if program[i] == 1 {
			reg_b = new(big.Int).Xor(reg_b, big.NewInt(int64(program[i+1])))
			//b ^= program[i+1]
		} else if program[i] == 2 {
			if program[i+1] < 4 {
				reg_b = new(big.Int).Mod(big.NewInt(int64(program[i+1])), big.NewInt(8))
				//b = program[i+1] % 8
			} else {
				if program[i+1] == 4 {
					reg_b = new(big.Int).Mod(reg_a, big.NewInt(8))
					//b = a % 8
				} else if program[i+1] == 5 {
					reg_b = new(big.Int).Mod(reg_b, big.NewInt(8))
					//b = b % 8
				} else if program[i+1] == 6 {
					reg_b = new(big.Int).Mod(reg_c, big.NewInt(8))
					//b = c % 8
				}
			}
		} else if program[i] == 3 {
			if reg_a.Cmp(big.NewInt(0)) == 0 {
				i += 2
				continue
			} else {
				i = program[i+1]
				continue
			}
		} else if program[i] == 4 {
			reg_b = new(big.Int).Xor(reg_b, reg_c)
			//b ^= c
		} else if program[i] == 5 {
			if program[i+1] < 4 {
				output = append(output, big.NewInt(int64(program[i+1]%8)))
			} else {
				if program[i+1] == 4 {
					output = append(output, new(big.Int).Mod(reg_a, big.NewInt(8)))
				} else if program[i+1] == 5 {
					output = append(output, new(big.Int).Mod(reg_b, big.NewInt(8)))
				} else if program[i+1] == 6 {
					output = append(output, new(big.Int).Mod(reg_c, big.NewInt(8)))
				}
			}
		} else if program[i] == 6 {
			if program[i+1] < 4 {
				reg_b = new(big.Int).Div(reg_a, new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(program[i+1])), nil))
				//b /= int(math.Pow(2, float64(program[i+1])))
			} else {
				if program[i+1] == 4 {
					reg_b = new(big.Int).Div(reg_a, new(big.Int).Exp(big.NewInt(2), reg_a, nil))
					//b /= int(math.Pow(2, float64(a)))
				} else if program[i+1] == 5 {
					reg_b = new(big.Int).Div(reg_a, new(big.Int).Exp(big.NewInt(2), reg_b, nil))
					//b /= int(math.Pow(2, float64(b)))
				} else if program[i+1] == 6 {
					reg_b = new(big.Int).Div(reg_a, new(big.Int).Exp(big.NewInt(2), reg_c, nil))
					//b /= int(math.Pow(2, float64(c)))
				}
			}
		} else if program[i] == 7 {
			if program[i+1] < 4 {
				reg_c = new(big.Int).Div(reg_a, new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(program[i+1])), nil))
				//c /= int(math.Pow(2, float64(program[i+1])))
			} else {
				if program[i+1] == 4 {
					reg_c = new(big.Int).Div(reg_a, new(big.Int).Exp(big.NewInt(2), reg_a, nil))
					//c /= int(math.Pow(2, float64(a)))
				} else if program[i+1] == 5 {
					reg_c = new(big.Int).Div(reg_a, new(big.Int).Exp(big.NewInt(2), reg_b, nil))
					//c /= int(math.Pow(2, float64(b)))
				} else if program[i+1] == 6 {
					reg_c = new(big.Int).Div(reg_a, new(big.Int).Exp(big.NewInt(2), reg_c, nil))
					//c /= int(math.Pow(2, float64(c)))
				}
			}
		}
		i += 2
	}

	fmt.Println(output)
	output_str := ""

	for i := range output {
		if i != len(output)-1 {
			output_str += fmt.Sprintf("%d,", output[i])
		} else {
			output_str += fmt.Sprintf("%d", output[i])
		}
	}

	out, err := os.Create("output.txt")
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("Part 1: " + output_str)
	shared.ErrCheck(err)
	out.Sync()
}

func GetData() (int, int, int, []int) {
	f, err := os.ReadFile("input.txt")
	shared.ErrCheck(err)

	rows := strings.Split(string(f), "\n")
	a, err := strconv.Atoi(strings.Split(rows[0], "Register A: ")[1])
	shared.ErrCheck(err)
	b, err := strconv.Atoi(strings.Split(rows[1], "Register B: ")[1])
	shared.ErrCheck(err)
	c, err := strconv.Atoi(strings.Split(rows[2], "Register C: ")[1])
	shared.ErrCheck(err)

	program_str := strings.Split(rows[4], "Program: ")[1]

	program := []int{}
	for i := range program_str {
		if program_str[i] != ',' {
			program = append(program, int(program_str[i])-48)
		}
	}

	return a, b, c, program
}
