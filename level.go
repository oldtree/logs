package logs

const (
	TraceLogLevel Level = iota
	InfoLogLevel
	NoticeLogLevel
	DebugLogLevel
	WarnLogLevel
	ErrorLogLevel
	PanicLogLevel
	AlertLogLevel
	FatalLogLevel
)

const (
	TraceLogLevelS  LevelS = "[TRACE]"
	InfoLogLevelS   LevelS = "[INFO]"
	NoticeLogLevelS LevelS = "[NOTICE]"
	DebugLogLevelS  LevelS = "[DEBUG]"
	WarnLogLevelS   LevelS = "[WARN]"
	ErrorLogLevelS  LevelS = "[ERROR]"
	PanicLogLevelS  LevelS = "[PANIC]"
	AlertLogLevelS  LevelS = "[ALERT]"
	FatalLogLevelS  LevelS = "[FATAL]"
)

const (
	DefaultTimeformat = "2006-01-02T15:04:05.000000000Z07:00"
	// time | file | line | func | level |real message
	TraceLogLevelFormat  = "%s %s %d %s " + TraceLogLevelS + " %s"
	InfoLogLevelFormat   = "%s %s %d %s " + InfoLogLevelS + " %s"
	NoticeLogLevelFormat = "%s %s %d %s " + NoticeLogLevelS + " %s"
	DebugLogLevelFormat  = "%s %s %d %s " + DebugLogLevelS + " %s"
	WarnLogLevelFormat   = "%s %s %d %s " + WarnLogLevelS + " %s"
	ErrorLogLevelFormat  = "%s %s %d %s " + ErrorLogLevelS + " %s"
	PanicLogLevelFormat  = "%s %s %d %s " + PanicLogLevelS + " %s"
	AlertLogLevelFormat  = "%s %s %d %s " + AlertLogLevelS + " %s"
	FatalLogLevelFormat  = "%s %s %d %s " + FatalLogLevelS + " %s"
)

type Level uint
type LevelS string

func (level Level) String() string {
	switch level {
	case DebugLogLevel:
		return string(DebugLogLevelS)
	case TraceLogLevel:
		return string(TraceLogLevelS)
	case InfoLogLevel:
		return string(InfoLogLevelS)
	case NoticeLogLevel:
		return string(NoticeLogLevelS)
	case WarnLogLevel:
		return string(WarnLogLevelS)
	case ErrorLogLevel:
		return string(ErrorLogLevelS)
	case PanicLogLevel:
		return string(PanicLogLevelS)
	case AlertLogLevel:
		return string(AlertLogLevelS)
	case FatalLogLevel:
		return string(FatalLogLevelS)
	}
	return string(DebugLogLevelS)
}

func (level Level) StringColor() string {
	switch level {
	case DebugLogLevel:
		return string(DebugLogLevelS)
	case TraceLogLevel:
		return string(TraceLogLevelS)
	case InfoLogLevel:
		return string(InfoLogLevelS)
	case NoticeLogLevel:
		return string(NoticeLogLevelS)
	case WarnLogLevel:
		return string(WarnLogLevelS)
	case ErrorLogLevel:
		return string(ErrorLogLevelS)
	case PanicLogLevel:
		return string(PanicLogLevelS)
	case AlertLogLevel:
		return string(AlertLogLevelS)
	case FatalLogLevel:
		return string(FatalLogLevelS)
	}
	return string(DebugLogLevelS)
}

func (levels LevelS) Level() uint {
	switch levels {
	case DebugLogLevelS:
		return uint(DebugLogLevel)
	case TraceLogLevelS:
		return uint(TraceLogLevel)
	case InfoLogLevelS:
		return uint(InfoLogLevel)
	case NoticeLogLevelS:
		return uint(NoticeLogLevel)
	case WarnLogLevelS:
		return uint(WarnLogLevel)
	case ErrorLogLevelS:
		return uint(ErrorLogLevel)
	case PanicLogLevelS:
		return uint(PanicLogLevel)
	case AlertLogLevelS:
		return uint(AlertLogLevel)
	case FatalLogLevelS:
		return uint(FatalLogLevel)
	}
	return uint(DebugLogLevel)
}
