package main

import (
	"context"
	"flag"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/kaneshin/spinner"
)

var interval float64
var duration float64
var command string

func init() {
	flag.Float64Var(&interval, "i", 0.05, "interval of spinner")
	flag.Float64Var(&duration, "d", 0.0, "duration of spinner")
	flag.StringVar(&command, "c", "", "commands are read from string")
	flag.Parse()
}

func run() error {
	var sp *spinner.Spinner
	var ctx context.Context
	var cancel context.CancelFunc
	if duration > 0.0 {
		d := time.Duration(1000*duration) * time.Millisecond
		ctx, cancel = context.WithTimeout(context.Background(), d)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel()
	ims := time.Duration(1000*interval) * time.Millisecond

	switch {
	case command != "":
		elm := strings.Fields(command)
		arg := elm[1:]
		sp = spinner.New(ims, func(ctx context.Context) {
			errCh := make(chan error)
			go func() {
				cmd := exec.CommandContext(ctx, elm[0], arg...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					errCh <- err
				}
				cancel()
			}()
			select {
			case <-ctx.Done():
				return
			case err := <-errCh:
				_ = err
				return
			}
		})
	default:
		sp = spinner.New(ims, func(ctx context.Context) {
			<-ctx.Done()
		})
	}

	sp.Do(ctx)

	return nil
}

func main() {
	err := run()
	if err != nil {
		os.Exit(1)
	}
}
