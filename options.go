package logruswr

type LogWrapperParams struct {
	formatter    Formatter
	output       Console
	level        Level
	logFilePath  string
	maxLogSize   int // in megabytes
	maxBackups   int
	maxAge       int // in days
	compressLogs bool
	hooks        []Hook
}

func newLogWrapperParams(opts ...Options) (*LogWrapperParams, error) {
	params := &LogWrapperParams{
		formatter: FormatterText,
		output:    Stdout,
		level:     Info,
	}

	for _, opt := range opts {
		if err := opt(params); err != nil {
			return nil, err
		}
	}

	return params, nil
}

func WithFormatter(f Formatter) Options {
	return func(p *LogWrapperParams) error {
		p.formatter = f
		return nil
	}
}

func WithOutput(c Console) Options {
	return func(p *LogWrapperParams) error {
		p.output = c
		return nil
	}
}

func WithLevel(l Level) Options {
	return func(p *LogWrapperParams) error {
		p.level = l
		return nil
	}
}

func WithLogFilePath(path string) Options {
	return func(p *LogWrapperParams) error {
		p.logFilePath = path
		return nil
	}
}

func WithMaxLogSize(size int) Options {
	return func(p *LogWrapperParams) error {
		p.maxLogSize = size
		return nil
	}
}

func WithMaxBackups(backups int) Options {
	return func(p *LogWrapperParams) error {
		p.maxBackups = backups
		return nil
	}
}

func WithMaxAge(age int) Options {
	return func(p *LogWrapperParams) error {
		p.maxAge = age
		return nil
	}
}

func WithCompressLogs(compress bool) Options {
	return func(p *LogWrapperParams) error {
		p.compressLogs = compress
		return nil
	}
}

func WithHooks(hooks ...Hook) Options {
	return func(p *LogWrapperParams) error {
		p.hooks = hooks
		return nil
	}
}

func WithHook(hook Hook) Options {
	return func(p *LogWrapperParams) error {
		p.hooks = append(p.hooks, hook)
		return nil
	}
}

// getters ....
func (p *LogWrapperParams) GetFormatter() Formatter {
	return p.formatter
}

func (p *LogWrapperParams) GetOutput() Console {
	return p.output
}

func (p *LogWrapperParams) GetLevel() Level {
	return p.level
}

func (p *LogWrapperParams) GetLogFilePath() string {
	return p.logFilePath
}

func (p *LogWrapperParams) GetMaxLogSize() int {
	return p.maxLogSize
}

func (p *LogWrapperParams) GetMaxBackups() int {
	return p.maxBackups
}

func (p *LogWrapperParams) GetMaxAge() int {
	return p.maxAge
}

func (p *LogWrapperParams) GetCompressLogs() bool {
	return p.compressLogs
}

func (p *LogWrapperParams) GetHooks() []Hook {
	return p.hooks
}

// setters ....

func (p *LogWrapperParams) SetFormatter(f Formatter) {
	p.formatter = f
}

func (p *LogWrapperParams) SetOutput(c Console) {
	p.output = c
}

func (p *LogWrapperParams) SetLevel(l Level) {
	p.level = l
}

func (p *LogWrapperParams) SetOptions(opts ...Options) error {
	for _, opt := range opts {
		if err := opt(p); err != nil {
			return err
		}
	}
	return nil
}

func (p *LogWrapperParams) SetLogFilePath(path string) {
	p.logFilePath = path
}

func (p *LogWrapperParams) SetMaxLogSize(size int) {
	p.maxLogSize = size
}

func (p *LogWrapperParams) SetMaxBackups(backups int) {
	p.maxBackups = backups
}

func (p *LogWrapperParams) SetMaxAge(age int) {
	p.maxAge = age
}

func (p *LogWrapperParams) SetCompressLogs(compress bool) {
	p.compressLogs = compress
}

func (p *LogWrapperParams) SetHooks(hooks ...Hook) {
	p.hooks = hooks
}
