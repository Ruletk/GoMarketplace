package logging

import (
	"github.com/sirupsen/logrus"
)

type LoggerNameHook struct {
	LoggerName string
}

func (hook *LoggerNameHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *LoggerNameHook) Fire(entry *logrus.Entry) error {
	entry.Data["loggerName"] = hook.LoggerName
	return nil
}
