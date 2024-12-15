package day14_24

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	shared "github.com/VladanT3/Advent_of_Code"
)

type coordinate struct {
	x int
	y int
}

type velocity struct {
	x int
	y int
}

type robot struct {
	pos coordinate
	vel velocity
}

func Part1() {
	robots := GetRobots()

	for range 100 {
		for _, robot := range robots {
			robot.pos.x += robot.vel.x
			robot.pos.y += robot.vel.y

			if robot.pos.x > 100 {
				robot.pos.x -= 100
			} else if robot.pos.x < 0 {
				robot.pos.x += 100
			}
			if robot.pos.y > 102 {
				robot.pos.y -= 102
			} else if robot.pos.y < 0 {
				robot.pos.y += 102
			}
		}
	}

	quadrant_1 := 0
	quadrant_2 := 0
	quadrant_3 := 0
	quadrant_4 := 0
	for _, robot := range robots {
		if robot.pos.x < 50 && robot.pos.y < 51 {
			quadrant_1++
		}
		if robot.pos.x > 50 && robot.pos.y < 51 {
			quadrant_2++
		}
		if robot.pos.x < 50 && robot.pos.y > 51 {
			quadrant_3++
		}
		if robot.pos.x > 50 && robot.pos.y > 51 {
			quadrant_4++
		}
	}

	safety_factor := quadrant_1 * quadrant_2 * quadrant_3 * quadrant_4
	fmt.Println(safety_factor)

	out, err := os.Create("output.txt")
	shared.ErrCheck(err)
	defer out.Close()

	_, err = out.WriteString("Part 1: " + strconv.Itoa(safety_factor))
	shared.ErrCheck(err)
	out.Sync()
}

func GetRobots() []robot {
	f, err := os.ReadFile("input.txt")
	shared.ErrCheck(err)

	out := []robot{}

	rows := strings.Split(string(f), "\n")
	for i := 0; i < len(rows)-1; i++ {
		row := strings.Split(rows[i], " ")
		pos := strings.Split(row[0], "=")
		vel := strings.Split(row[1], "=")
		pos = strings.Split(pos[1], ",")
		vel = strings.Split(vel[1], ",")

		x, err := strconv.Atoi(pos[0])
		shared.ErrCheck(err)
		y, err := strconv.Atoi(pos[1])
		shared.ErrCheck(err)
		cord := coordinate{x, y}

		vel_x, err := strconv.Atoi(vel[0])
		shared.ErrCheck(err)
		vel_y, err := strconv.Atoi(vel[1])
		shared.ErrCheck(err)
		vels := velocity{vel_x, vel_y}

		out = append(out, robot{cord, vels})
	}

	return out
}
