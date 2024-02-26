package hooks

import "github.com/sirupsen/logrus"

type TestHook struct {
	Fired bool
}

func (h *TestHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *TestHook) Fire(entry *logrus.Entry) error {
	h.Fired = true
	return nil
}

func (h *TestHook) Reset() {
	h.Fired = false
}

func (h *TestHook) IsFired() bool {
	return h.Fired
}
