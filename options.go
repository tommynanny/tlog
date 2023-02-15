package tlog

import "time"

func Config(fns ...OptionFunc) { WithOptions(fns...) }
func WithOptions(fns ...OptionFunc) {
	for _, fn := range fns {
		fn(GlobalOptions)
	}
}

// Options for dumper
type Options struct {
	// PrintToStdout determines whether log messages are printed to standard output in addition to the log file.
	AutoSetupLogger bool
	LOG_ROOT        string
}

// OptionFunc type
type OptionFunc func(opts *Options)

// NewDefaultOptions create.
func NewDefaultOptions() *Options {
	return &Options{
		AutoSetupLogger: true,
		LOG_ROOT:        "logs/" + time.Now().UTC().Format("2006-01-02--15-04-05"),
	}
}

func NoAutoSetup() OptionFunc {
	return func(opt *Options) {
		opt.AutoSetupLogger = false
	}
}

func SetLOG_ROOT(root string) OptionFunc {
	return func(opt *Options) {
		opt.LOG_ROOT = root
	}
}
