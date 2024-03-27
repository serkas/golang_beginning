package main

import (
	"context"
	"os/signal"
	"syscall"
)

func main() {
	// Scenario: We need to run sum process and then stop it at some event
	//runStoppingWithChannel()

	//runStoppingWithContext()

	// graceful shutdown example
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM) // this will intercept OS signals (like Ctrl+C)
	defer stop()

	runWithExternalContext(ctx)
}
