package server

import (
	"os"
)

type StartRunner struct {
	ListenAddr string
	Server     *GracefulShutdown // Change this to store the GracefulShutdown instance
}

func (runner *StartRunner) Run() error {
	server := &GracefulShutdown{
		ListenAddr: runner.ListenAddr,
	}
	runner.Server = server // Assign the server instance to the StartRunner.Server field
	server.Start()
	return nil
}

// Shutdown is a Graceful shutdown method (update method name to match GracefulShutdown struct)
func (runner *StartRunner) Shutdown() {
	runner.Server.Shutdown() // Call the Shutdown method of the GracefulShutdown instance
	os.Exit(0)
}
