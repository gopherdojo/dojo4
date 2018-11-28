package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func isDone(t *testing.T, ctx context.Context) bool {
	t.Helper()
	select {
	case <-ctx.Done():
		return true
	case <-time.After(1 * time.Second):
		return false
	}
}

func Test_sigHandledContext(t *testing.T) {
	var tests = []struct {
		name     string
		canceled bool
		sig      syscall.Signal
	}{
		{"cancel context if send SIGHUP", true, syscall.SIGHUP},
		{"cancel context if send SIGINT", true, syscall.SIGINT},
		{"cancel context if send SIGTERM", true, syscall.SIGTERM},
		{"cancel context if send SIGQUIT", true, syscall.SIGQUIT},
		{"not cancel context if send SIGWINCH", false, syscall.SIGWINCH},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ch := make(chan os.Signal, 1)
			defer signal.Stop(ch)
			signal.Notify(ch,
				syscall.SIGHUP,
				syscall.SIGINT,
				syscall.SIGTERM,
				syscall.SIGQUIT)

			ctx, cancel := sigHandledContext(ch)
			defer cancel()
			if err := syscall.Kill(syscall.Getpid(), tt.sig); err != nil {
				t.Fatalf("Failed to send SIGNAL %v to this process", tt.sig)
			}

			got := isDone(t, ctx)
			if got != tt.canceled {
				t.Errorf("Channel trapping signal cancel or not the context: got %v, want %v", got, tt.canceled)
			}
		})
	}
}
