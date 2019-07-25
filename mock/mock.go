package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

// Sleeper inTerFace
type Sleeper interface {
	Sleep(d time.Duration)
}

// DefaultSleeper struct
type DefaultSleeper struct {
	duration time.Duration
}

// Sleep Defaulet
func (ds DefaultSleeper) Sleep(d time.Duration) {
	time.Sleep(ds.duration)
}

func main() {
	Countdown(os.Stdout, &DefaultSleeper{}, 1*time.Second)
}

// Countdown type
func Countdown(writer io.Writer, sleeper Sleeper, sleepTime time.Duration) {
	for i := 3; i > 0; i-- {
		sleeper.Sleep(sleepTime)
		fmt.Fprintln(writer, strconv.Itoa(i))
	}
	sleeper.Sleep(sleepTime)
	fmt.Fprintf(writer, "Go!")
}
