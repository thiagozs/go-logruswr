package logruswr

import "github.com/sirupsen/logrus"

type Options func(*LogWrapperParams) error

type Level uint32

type Console int

type Formatter int

type Entry = logrus.Entry

type Hook = logrus.Hook

type Fields map[string]interface{}
