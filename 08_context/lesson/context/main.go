package main

func main() {
	runStoppingWithChannel()

	//runStoppingWithContext()

	// graceful shutdown example
	//ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM) // this will intercept OS signals (like Ctrl+C)
	//defer stop()
	//
	//runWithExternalContext(ctx)
}
