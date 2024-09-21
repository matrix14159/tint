package tint

import (
	"context"
	"log/slog"
	"runtime"
	"time"
)

func (h *handler) MarkError(msg string, a ...any) error {
	var pc uintptr

	var pcs [1]uintptr
	// skip [runtime.Callers, this function, this function's caller]
	runtime.Callers(3, pcs[:])
	pc = pcs[0]

	r := slog.NewRecord(time.Now(), slog.LevelError, msg, pc)
	r.Add(a...)
	return h.Handle(context.Background(), r)
}
