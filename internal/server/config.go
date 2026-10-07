package server

import "github.com/roemer/scubanaut/internal/store"

type Config struct {
	Address string
	Version string
	Store   store.Client
}
