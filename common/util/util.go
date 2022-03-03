package util

import (
	"os"

	"github.com/go-kratos/kratos/v2/log"
)

func NewTestLogger() log.Logger {
	return log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.Caller(4),
	)
}
