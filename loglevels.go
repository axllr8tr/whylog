package whylog
// Logging levels and their properties

var DefaultLevel = InfoLevel

type LogLevel struct {
	Emoji	    string
	Name	    string
	Syslog	    int
	Priority    int
	Style       LogStyle
}

type LogStyle struct {
	DateColor    string
	LevelColor   string
	MessageColor string
}

// Default styles
var (
	traceStyle  = LogStyle { "\033[30;1m", "\033[30;1m", "\033[30;1m" }
	debugStyle  = LogStyle { "\033[30;1m", "\033[34m", "\033[30;1m" }
	infoStyle   = LogStyle { "\033[30;1m", "\033[34;1m", "\033[37m" }
	noticeStyle = LogStyle { "\033[30;1m", "\033[36;1m", "\033[37m" }
	warnStyle   = LogStyle { "\033[33;1m", "\033[33;1m", "\033[37m" }
	errorStyle  = LogStyle { "\033[31m", "\033[31m", "\033[37m" }
	critStyle   = LogStyle { "\033[31;1m", "\033[31;1m", "\033[37m" }
	alertStyle  = LogStyle { "\033[31;1m", "\033[31;1;5m", "\033[37;1m" }
	emergStyle  = LogStyle { "\033[37;41;1m", "\033[37;41;1;5m", "\033[37;41;1m" }
)

// Custom pre-defined styles
var (
	GrayStyle    = LogStyle { "\033[30;1m", "\033[30;1m", "\033[37m" }
	RedStyle     = LogStyle { "\033[30;1m", "\033[31;1m", "\033[37m" }
	GreenStyle   = LogStyle { "\033[30;1m", "\033[32;1m", "\033[37m" }
	YellowStyle  = LogStyle { "\033[30;1m", "\033[33;1m", "\033[37m" }
	BlueStyle    = LogStyle { "\033[30;1m", "\033[34;1m", "\033[37m" }
	MagentaStyle = LogStyle { "\033[30;1m", "\033[35;1m", "\033[37m" }
	CyanStyle    = LogStyle { "\033[30;1m", "\033[36;1m", "\033[37m" }
	WhiteStyle   = LogStyle { "\033[30;1m", "\033[37;1m", "\033[37m" }
)

// Default available loglevels
var (
	// Granular information about function calls
	TraceLevel = LogLevel {
		"\U00002699\uFE0F", // cog + emoji variant
		"TRACE",
		7,
		20,
		traceStyle,
	}

	// Information useful for debugging, e.g. API calls, performance metrics, queue statuses 
	DebugLevel = LogLevel {
		"\U0001F41B", // bug
		"DEBUG",
		7,
		30,
		debugStyle,
	}

	// Basic operational messages, anything more verbose is not shown by default
	InfoLevel = LogLevel {
		"\U00002139\uFE0F", // info [i] + emoji variant
		"INFO",
		6,
		40,
		infoStyle,
	}

	// Significant non-error events, e.g. backups, synchronization, account creation (if running a service)
	NoticeLevel = LogLevel {
		"\U0001F440", // eyes
		"NOTICE",
		5,
		50,
		noticeStyle,
	}
	
	// Issues that may cause the service to function not as well as expected, e.g. performance issues, usage of deprecated APIs
	WarnLevel = LogLevel {
		"\U000026A0\uFE0F", // warning + emoji variant
		"WARNING",
		4,
		60,
		warnStyle,
	}

	// Errors that may seriously mess with the program's operation, e.g. broken configs, database connection faults
	ErrorLevel = LogLevel {
		"\U0000274C", // x
		"ERROR",
		3,
		70,
		errorStyle,
	}

	// Really serious errors, e.g. data integrity issues
	CritLevel = LogLevel {
		"\U000026D4", // no entry
		"CRITICAL",
		2,
		80,
		critStyle,
	}

	// Errors that require immediate action, e.g. compute resource exhaustion
	AlertLevel = LogLevel {
		"\U0001F6A8", // emergency signal
		"ALERT",
		1,
		90,
		alertStyle,
	}

	// Errors the program sends out right before quitting
	EmergLevel = LogLevel {
		"\U0001F4A3", // bomb,
		"EMERGENCY",
		0,
		100,
		emergStyle,
	}
)

// Returns a custom loglevel with same priority as DEBUG
func NewGenericLevel(name string, emoji string) LogLevel {
	return LogLevel {
		emoji,
		name,
		7,
		30,
		MagentaStyle,
	}
}
