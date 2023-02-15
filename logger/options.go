package logger

func (tl *TLogger) Config(fns ...OptionFunc) { tl.WithOptions(fns...) }
func (t *TLogger) WithOptions(fns ...OptionFunc) *TLogger {
	for _, fn := range fns {
		fn(t.Options)
	}
	return t
}

// Options for dumper
type LoggerOptions struct {
	// PrintToStdout determines whether log messages are printed to standard output in addition to the log file.
	PrintToStdout bool
	// UseWrapper determines whether a log file wrapper is used.
	UseWrapper bool
	// ColorfulStdout determines whether colorful output is used for standard output.
	ColorfulStdout bool
	// WithCallerSkip is the number of stack frames to skip when determining the caller's location.
	WithCallerSkip int
}

// OptionFunc type
type OptionFunc func(opts *LoggerOptions)

// NewDefaultOptions create.
func NewDefaultOptions() *LoggerOptions {
	return &LoggerOptions{
		UseWrapper:     true,
		ColorfulStdout: true,
		PrintToStdout:  true,
		WithCallerSkip: 4,
	}
}

func NoStdout() OptionFunc {
	return func(opt *LoggerOptions) {
		opt.PrintToStdout = false
	}
}

func NoWrapper() OptionFunc {
	return func(opt *LoggerOptions) {
		opt.UseWrapper = false
	}
}

func NoColorfulStdout() OptionFunc {
	return func(opt *LoggerOptions) {
		opt.ColorfulStdout = false
	}
}

func WithCallSkip(skip int) OptionFunc {
	return func(opt *LoggerOptions) {
		opt.WithCallerSkip = skip
	}
}
