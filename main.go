package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	screen.Clear()

	for {
		t := time.Now()
		screen.MoveTopLeft()

		h, m, s := t.Hour(), t.Minute(), t.Second()
		clock := CreateClock(h, m, s)

		for line := range clock[0] {
			for i, digit := range clock {
				next := clock[i][line]
				if digit == clock[2] && s%2 == 0 {
					next = "  "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}
		time.Sleep(time.Second)
	}
}
