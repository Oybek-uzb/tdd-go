package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	sleeper := &DefaultSleeper{}
	CountDown(os.Stdout, sleeper)
}

const countDownStart = 3
const finalWord = "Go!"

func CountDown(writer io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}

	sleeper.Sleep()
	fmt.Fprint(writer, finalWord)
}

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (ds *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
