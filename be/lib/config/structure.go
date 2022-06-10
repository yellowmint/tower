package config

import (
	"git.jetbrains.space/artdecoction/wt/tower/lib/fbase"
	"git.jetbrains.space/artdecoction/wt/tower/lib/logs"
)

type mode string

const (
	Dev  mode = "dev"
	Prod      = "prod"
)

type LoggerFormat string

type Config struct {
	Version                   string
	Port                      string
	LoggerFormat              logs.Format
	GcpProjectId              string
	FirebaseEmulator          fbase.EmulatorConfig
	AuthenticationMockEnabled bool
}
