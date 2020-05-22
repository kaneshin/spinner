package spinner

import (
	"context"
	"fmt"
	"time"
)

type Spinner struct {
	elements []string
	interval time.Duration
	run      func(context.Context)
	wait     chan int
}

func New(interval time.Duration, run func(context.Context)) *Spinner {
	return &Spinner{
		elements: []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		run:      run,
		interval: interval,
		wait:     make(chan int),
	}
}

func (s *Spinner) start() {
	elements := s.elements
	for {
		select {
		case <-s.wait:
			return
		default:
			for _, v := range elements {
				fmt.Printf("  %s  \x0d", v)
				time.Sleep(s.interval)
			}
		}
	}
}

func (s *Spinner) Do(ctx context.Context) {
	go s.start()
	s.run(ctx)
	s.wait <- 0
}

func (s *Spinner) Parallel() {
	// TODO: parallelism
	return
}

func (s *Spinner) Wait() {
	// TODO: waiting proc in parallel
	return
}

func Run(ctx context.Context, interval time.Duration, run func(*Spinner) func(context.Context)) *Spinner {
	sp := New(interval, nil)
	sp.run = run(sp)
	sp.Do(ctx)
	return sp
}
