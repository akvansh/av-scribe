package scribe

import (
	"log/slog"
)

type Level int8

const (
	LevelDebug Level = Level(slog.LevelDebug)
	LevelInfo Level = Level(slog.LevelInfo)
	LevelWarn Level = Level(slog.LevelWarn)
	LevelError Level = Level(slog.LevelError)
)

func (l Level) String() string {
	switch l {

	}
}
