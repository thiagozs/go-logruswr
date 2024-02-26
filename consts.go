package logruswr

import "fmt"

const (
	Panic Level = iota
	Fatal
	Error
	Warn
	Info
	Debug
	Trace

	Stdout Console = iota
	Stderr
	File

	FormatterText Formatter = iota
	FormatterJSON
)

func (c Console) String() string {
	switch c {
	case Stdout:
		return "stdout"
	case Stderr:
		return "stderr"
	default:
		return "stdout"
	}
}

func (l Level) String() string {
	switch l {
	case Panic:
		return "panic"
	case Fatal:
		return "fatal"
	case Error:
		return "error"
	case Warn:
		return "warn"
	case Info:
		return "info"
	case Debug:
		return "debug"
	case Trace:
		return "trace"
	default:
		return "info"
	}
}

func (l Level) MarshalText() ([]byte, error) {
	return []byte(l.String()), nil
}

func (l *Level) UnmarshalText(text []byte) error {
	switch string(text) {
	case "panic":
		*l = Panic
	case "fatal":
		*l = Fatal
	case "error":
		*l = Error
	case "warn":
		*l = Warn
	case "info":
		*l = Info
	case "debug":
		*l = Debug
	case "trace":
		*l = Trace
	default:
		return fmt.Errorf("unknown level: %s", text)
	}
	return nil
}

func (f Formatter) String() string {
	switch f {
	case FormatterText:
		return "text"
	case FormatterJSON:
		return "json"
	default:
		return "text"
	}
}
