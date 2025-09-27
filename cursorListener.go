package main

import (
	"fmt"
	"time"
	"flag"
	"os/exec"

	hook "github.com/robotn/gohook"
)

var nextTime time.Time

func main(){
	chanHook := hook.Start()
	defer hook.End()

	delta_time := flag.Int("delta-time", 5, "Время бездействия пользователя в минутах")
	browser_name := flag.String("browser-name", "chromium-browser", "Имя браузера, который будет убит по истечению времени бездействия")

	flag.Usage = func() {
		fmt.Println("Программа для убийства браузера при бездействии")
		fmt.Println("\nВерсия 1.0.2 от 25.09.2025")
		fmt.Println("\t- Раньше это программа убивала компьютеры")
		fmt.Println("\nРазработана ГОКУ ЦЗН Мурманской Области инспектором ЦЗН")
		fmt.Println("Глущенко Евгением Юрьевичем\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	now := time.Now()

	nextTime = now.Add(time.Minute * time.Duration(*delta_time))

	for {
		select {
		case <-chanHook:
			now := time.Now()
			nextTime = now.Add(time.Minute * time.Duration(*delta_time))
		default:
			now := time.Now()

			if now.Hour() >= nextTime.Hour() &&
				now.Minute() >= nextTime.Minute() &&
				now.Second() >= nextTime.Second() {
				cmd := exec.Command("pkill", "-f", *browser_name)
				cmd.Start()
				nextTime = now.Add(time.Minute * time.Duration(*delta_time))
			}	
		}
	}
}
