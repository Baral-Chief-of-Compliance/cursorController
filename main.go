package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/go-vgo/robotgo"
)

const DELTA_MINUTE int = 5
const BOWSER_NAME string = "chromium"

func main() {
	now := time.Now()

	nextTime := now.Add(time.Minute * time.Duration(DELTA_MINUTE))

	startX, startY := robotgo.Location()
	fmt.Println("pos: ", startX, startY)

	for {
		now := time.Now()

		newX, newY := robotgo.Location()

		fmt.Println("pos: ", newX, newY)
		if newX != startX && newY != startY {
			startX = newX
			startY = newY
			fmt.Println("Таймер обнулился")
			nextTime = now.Add(time.Minute * time.Duration(DELTA_MINUTE))
		}

		if now.Hour() >= nextTime.Hour() &&
			now.Minute() >= nextTime.Minute() &&
			now.Second() >= nextTime.Second() {
			fmt.Println("Время достигнуто")
			cmd := exec.Command("sudo", "pkill", "-f", "chromium-browser")
			cmd.Start()
		}

	}
}
