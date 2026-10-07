package main

import "github.com/roemer/scubanaut/internal/server"

// Default if not build via CI and the -ldflags
var version = "dev"

func main() {
	config := server.Config{
		Address: ":8080",
		Version: version,
	}
	// Start HTTP server
	srv := server.New(config)
	srv.Start()
}
