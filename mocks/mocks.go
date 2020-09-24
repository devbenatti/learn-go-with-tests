package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const lastWord = "Go!"
const beginCount = 3

type sleeper interface {
	Sleep()
}

func score(writer io.Writer, sleeper sleeper) {
	for i := beginCount; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, fmt.Sprint(i))
	}
	sleeper.Sleep()
	fmt.Fprintln(writer, lastWord)
}

type sleeperConfig struct {
	duration time.Duration
	pause    func(time.Duration)
}

func (s *sleeperConfig) Sleep() {
	s.pause(s.duration)
}

func main() {
	sleeper := &sleeperConfig{2 * time.Second, time.Sleep}
	score(os.Stdout, sleeper)
}
