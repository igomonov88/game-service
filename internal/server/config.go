package server

import "time"

type Config struct {
	Address         string
	ReadTimeout     time.Duration
	WriteTimout     time.Duration
	ShutdownTimeout time.Duration
}
