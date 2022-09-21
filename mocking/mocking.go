package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

func main() {
	log.Println("Program started.")
	Countdown(os.Stdout)
}

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(out, finalWord)

}
