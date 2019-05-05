package test

import (
	"github.com/1-bi/clog/loggercom"
	"github.com/1-bi/clog/loggerzap"
	"github.com/1-bi/clog/loggerzap/appender"
	zaplayout "github.com/1-bi/clog/loggerzap/layout"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
)

//  Test_BasicCase1_Debug define bug info
func Test_Zap_Factory_case1_advanced(t *testing.T) {

	var lfm loggercom.LoggerFactory
	var lfo = loggerzap.NewLoggerOption()
	lfo.SetLevel("warn")

	// use new or struct binding
	// create instance from implement
	lfm = loggercom.NewLoggerFactory(new(loggerzap.ZapFactoryRegister), lfo)

	// --- create logger factory manager
	if lfm == nil {
		t.Errorf(": logger factory  expected,[%v], actually: [%v]", " object ", " is null ")
	}

}

//  Test_BasicCase1_Debug define bug info
func Test_Zap_Factory_case1_advanced_example(t *testing.T) {

	var lfm loggercom.LoggerFactory

	var multiOpts []loggercom.Option
	multiOpts = make([]loggercom.Option, 0)

	// --- construct layout ---
	var jsonLayout = zaplayout.NewJsonLayout()
	//jsonLayout.SetTimeFormat("2006-01-02 15:04:05")
	jsonLayout.SetTimeFormat("2006-01-02 15:04:05 +0800 CST")
	//fmt.Println( time.Now().Location() )

	// --- set appender
	var consoleAppender = appender.NewConsoleAppender(jsonLayout)

	var loggerOpt1 = loggerzap.NewLoggerOption()
	loggerOpt1.SetLevel("debug")
	loggerOpt1.AddAppender(consoleAppender)

	multiOpts = append(multiOpts, loggerOpt1)

	var fileAppender = appender.NewFileAppender(jsonLayout)

	var loggerOpt2 = loggerzap.NewLoggerOption()
	loggerOpt2.SetLevel("warn")
	loggerOpt2.AddAppender(fileAppender)

	//multiOpts = append(multiOpts, loggerOpt2)

	// use new or struct binding
	// create instance from implement
	lfm = loggercom.NewLoggerFactory(new(loggerzap.ZapFactoryRegister), multiOpts...)

	// --- create logger factory manager
	if lfm == nil {
		t.Errorf(": logger factory  expected,[%v], actually: [%v]", " object ", " is null ")
	}

	//logger := lfm.GetLogger()
	logger := loggercom.GetLogger("module")

	logger.Debug("debug message for  example", nil)
	logger.Info("info message for  example", nil)
	logger.Warn("warn message for  example", nil)
	logger.Error("error  message for  example", nil)

}

func Test_Zap_Factory_anothe(t *testing.T) {

	// First, define our level-handling logic.
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})
	//lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	//	return lvl < zapcore.ErrorLevel
	//})

	// Assume that we have clients for two Kafka topics. The clients implement
	// zapcore.WriteSyncer and are safe for concurrent use. (If they only
	// implement io.Writer, we can use zapcore.AddSync to add a no-op Sync
	// method. If they're not safe for concurrent use, we can add a protecting
	// mutex with zapcore.Lock.)
	//topicDebugging := zapcore.AddSync(ioutil.Discard)
	//topicErrors := zapcore.AddSync(ioutil.Discard)

	// High-priority output should also go to standard error, and low-priority
	// output should also go to standard out.
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	// Optimize the Kafka output for machine consumption and the console output
	// for human operators.
	//kafkaEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	core := zapcore.NewTee(
		//zapcore.NewCore(kafkaEncoder, topicErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		//zapcore.NewCore(kafkaEncoder, topicDebugging, lowPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, highPriority),
	)

	// From a zapcore.Core, it's easy to construct a Logger.
	logger := zap.New(core)
	defer logger.Sync()
	logger.Info("constructed a logger")
	logger.Warn("constructed a logger waring ")
	logger.Error("constructed a logger error ")

}
