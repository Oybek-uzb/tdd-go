package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	CountDown(os.Stdout)
}

const countDownStart = 3
const finalWord = "Go!"

func CountDown(writer io.Writer) {
	for i := countDownStart; i >= 0; i-- {
		fmt.Fprintln(writer, i)
	}

	fmt.Fprint(writer, finalWord)
}
