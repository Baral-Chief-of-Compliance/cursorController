package main

import (
	"flag"
	"fmt"
	"os/exec"
	"time"

	"github.com/go-vgo/robotgo"
)

var nextTime time.Time

func main() {
	delta_time := flag.Int("delta-time", 5, "Время бездействия пользователя в минутах")
	browser_name := flag.String("browser-name", "chromium-browser", "Имя браузера, который будет убит по истечению времени бездействия")

	flag.Usage = func() {
		fmt.Println("Программа для убийства браузера при бездействии")
		fmt.Println("\nВерсия 1.0.1 от 15.08.2025")
		fmt.Println("\t- Раньше это программа убивала компьютеры")
		fmt.Println("\nРазработана ГОКУ ЦЗН Мурманской Области инспектором ЦЗН")
		fmt.Println("Глущенко Евгением Юрьевичем\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	now := time.Now()

	nextTime = now.Add(time.Minute * time.Duration(*delta_time))

	startX, startY := robotgo.Location()

	for {
		now := time.Now()

		newX, newY := robotgo.Location()

		if newX != startX && newY != startY {
			startX = newX
			startY = newY
			nextTime = now.Add(time.Minute * time.Duration(*delta_time))
		}

		if now.Hour() >= nextTime.Hour() &&
			now.Minute() >= nextTime.Minute() &&
			now.Second() >= nextTime.Second() {
			cmd := exec.Command("pkill", "-f", *browser_name)
			cmd.Start()
			nextTime = now.Add(time.Minute * time.Duration(*delta_time))
		}

	}
}
