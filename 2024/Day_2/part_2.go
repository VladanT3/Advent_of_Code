package day2

import (
	"os"
	"strconv"
)

func Part2() {
	data := GetData()
	num_of_safe := 0
	var increasing bool
	var direction_set bool
	safe := false
	dampened := false

	for _, arr := range data {
		for i := 0; i < len(arr)-1; i++ {
			if !direction_set {
				if arr[i] > arr[i+1] {
					increasing = false
				} else if arr[i] < arr[i+1] {
					increasing = true
				} else {
					if !dampened {
						dampened = true
						continue
					} else {
						safe = false
						break
					}
				}
				direction_set = true
			}

			if increasing {
				if arr[i] >= arr[i+1] {
					if !dampened {
						dampened = true
						continue
					} else {
						safe = false
						break
					}
				}
				if arr[i]-arr[i+1] < -3 || arr[i]-arr[i+1] > -1 {
					if !dampened {
						dampened = true
						continue
					} else {
						safe = false
						break
					}
				}
			} else {
				if arr[i] <= arr[i+1] {
					if !dampened {
						dampened = true
						continue
					} else {
						safe = false
						break
					}
				}
				if arr[i]-arr[i+1] > 3 || arr[i]-arr[i+1] < 1 {
					if !dampened {
						dampened = true
						continue
					} else {
						safe = false
						break
					}
				}
			}
			safe = true
		}
		if safe {
			num_of_safe++
			safe = false
		}
		direction_set = false
		dampened = false
	}

	f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	ErrCheck(err)
	defer f.Close()

	_, err = f.WriteString("\nPart 2: " + strconv.Itoa(num_of_safe))
	ErrCheck(err)
	f.Sync()
}
