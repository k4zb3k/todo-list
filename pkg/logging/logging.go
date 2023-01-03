package logging

import (
	"io"

	"github.com/sirupsen/logrus"
)

type writerHook struct {
	Writer []io.Writer
	LogLevels []logrus.Level
}

