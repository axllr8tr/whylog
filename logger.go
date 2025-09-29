package whylog

import (
	"fmt"
	"os"
	"sync"
	"time"

	"golang.org/x/term"
)

type Logger struct {
	MinLevel    LogLevel
	Emoji	    string
	Name	    string
	File	    *os.File
	EnableColor bool
	EnableEmoji bool
	IsATTY	    bool
	mu	    sync.Mutex
	formatter   func(ctx *Logger, useColors bool, useEmoji bool, stamp time.Time, level LogLevel, message string, args ...any) string
}

const noColor = "\033[0m"

// Simple logger suited for development environments

func simpleLoggerFmt(ctx *Logger, useColors bool, useEmoji bool, stamp time.Time, level LogLevel, message string, av ...any) string {
	prefix := ""

	timestamp := stamp.Format("2006-01-02 15:04:05.000000000")
	loglevel := level.Name

	msg := message
	if len(av) != 0 {
		msg += " " + fmt.Sprintf("%v", av)
	}
	
	if useEmoji {
		sep := " "
		if ctx.IsATTY {
			sep = "\u200B"
		}
		if ctx.Emoji != "" {
			prefix = ctx.Emoji
		}
		if level.Emoji != "" {
			loglevel = level.Emoji + sep + loglevel
		}
	}

	if useColors {
		timestamp = level.Style.DateColor + timestamp + noColor
		loglevel = level.Style.LevelColor + loglevel + noColor
		msg = level.Style.MessageColor + msg + noColor
	}
	return fmt.Sprintf("%s\t%s\t%s\t	%s\t\t%s", prefix, timestamp, ctx.Name, loglevel, msg)
}

func NewSimpleLogger() *Logger {
	return &Logger {
		MinLevel:    DefaultLevel,
		Emoji:       "",
		Name:        "",
		File:        os.Stderr,
		EnableColor: true,
		EnableEmoji: true,
		IsATTY:      term.IsTerminal(int(os.Stderr.Fd())), // Will disable colors if ends up being false
		formatter:   simpleLoggerFmt,
	}
}


//////

func (l *Logger) Log(lvl LogLevel, message string, av ...any) {
	if (lvl.Priority < l.MinLevel.Priority) {
		return
	}
	fmt.Fprint(l.File, l.formatter(
		l,
		l.EnableColor && l.IsATTY,
		l.EnableEmoji,
		time.Now(),
		lvl,
		message,
		av...,
	), "\n")
}

func (l *Logger) Logf(lvl LogLevel, format string, av ...any) {
	l.Log(lvl, fmt.Sprintf(format, av...))
}
