package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	CountDown(os.Stdout)
}

func CountDown(writer io.Writer) {
	fmt.Fprint(writer, "3")
}
