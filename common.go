package logzap

import (
	loggercom "github.com/1-bi/log-api"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

const (

	// --- key contant ---
	PRESETS = "zap.presets"

	// --- value constant ---
	PRESETS_EXAMPLE = "example"
	PRESETS_PROD    = "prod"
	PRESETS_DEV     = "dev"
	PRESETS_NOP     = "nop"
)

func levelEventFilter(runtimeLevel byte) zapcore.LevelEnabler {

	var levelEnabler zapcore.LevelEnabler

	// set the runtime level
	/*
		switch runtimeLevel {
		case loggercom.DEVEL_DEBUG:

			// return define level
			levelEnabler = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl >= zapcore.DebugLevel
			})
			break

		case loggercom.DEVEL_INFO:
			levelEnabler = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl >= zapcore.InfoLevel
			})
			break

		case loggercom.DEVEL_WARN:

			levelEnabler = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl >= zapcore.WarnLevel
			})
			break
		case loggercom.DEVEL_FATAL:

			levelEnabler = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl >= zapcore.FatalLevel
			})
			break

		case loggercom.DEVEL_ERROR:

			levelEnabler = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl >= zapcore.ErrorLevel
			})
			break
		}

		if levelEnabler == nil {

			levelEnabler = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl >= zapcore.PanicLevel
			})

		}
	*/

	switch runtimeLevel {
	case loggercom.DEBUG:

		// return define level
		levelEnabler = zap.NewAtomicLevelAt(zap.DebugLevel).Level()
	case loggercom.INFO:
		levelEnabler = zap.NewAtomicLevelAt(zap.InfoLevel).Level()
	case loggercom.WARN:
		levelEnabler = zap.NewAtomicLevelAt(zap.WarnLevel).Level()
	case loggercom.FATAL:
		levelEnabler = zap.NewAtomicLevelAt(zap.FatalLevel).Level()
	case loggercom.ERROR:
		levelEnabler = zap.NewAtomicLevelAt(zap.ErrorLevel).Level()
	}

	if levelEnabler == nil {
		levelEnabler = zap.NewAtomicLevelAt(zap.PanicLevel).Level()
	}

	return levelEnabler
}

// timeformater define thime formater
type TimeFormater struct {
	pattern    string
	timezoneId string
}

func NewTimeFormater(pattern string) *TimeFormater {
	tf := new(TimeFormater)
	tf.pattern = pattern
	return tf
}

func (myself *TimeFormater) SetPattern(newPattern string) {
	myself.pattern = newPattern

	// --- parse fpr itc
}

func (myself *TimeFormater) SetTimeZone(newTimezone string) {
	myself.timezoneId = newTimezone
}

func (myself *TimeFormater) CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {

	var tmpTime time.Time
	if myself.timezoneId == "" || myself.timezoneId == "UTC" {
		tmpTime = t.UTC()
	} else {
		tmpTime = t
	}

	if myself.pattern == "" {
		myself.pattern = "2006-01-02T15:04:05.000Z0700"
	}

	enc.AppendString(tmpTime.Format(myself.pattern))
}
