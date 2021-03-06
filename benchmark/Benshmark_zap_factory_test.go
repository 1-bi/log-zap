package benchmark

import (
	logapi "github.com/1-bi/log-api"
	loggerzap "github.com/1-bi/log-zap"
	appender "github.com/1-bi/log-zap/appender"
	zaplayout "github.com/1-bi/log-zap/layout"
	"log"
	"testing"
)

//  Test_BasicCase1_Debug define bug info
func Benchmark_Zap_Factory_case1_advanced_example(b *testing.B) {
	//b.StopTimer()
	var multiOpts = make([]logapi.Option, 0)
	// --- construct layout ---
	var jsonLayout = zaplayout.NewJsonLayout()
	// --- set appender
	var consoleAppender = appender.NewConsoleAppender(jsonLayout)

	var loggerOpt1 = loggerzap.NewLoggerOption()
	loggerOpt1.SetLevel("info")
	loggerOpt1.AddAppender(consoleAppender)
	multiOpts = append(multiOpts, loggerOpt1)

	jsonLayout = zaplayout.NewJsonLayout()
	//jsonLayout.SetTimeFormat("2006-01-02 15:04:05")
	jsonLayout.SetTimeFormat("2006-01-02T15:04:05.0700Z ")
	//jsonLayout.SetTimezoneId("UTC")

	var specOpt1 = loggerzap.NewLoggerOption()
	specOpt1.SetLoggerPattern("benshmark.test")
	specOpt1.SetLevel("warn")
	specOpt1.AddAppender(appender.NewConsoleAppender(jsonLayout))
	multiOpts = append(multiOpts, loggerOpt1)

	// use new or struct binding
	// create instance from implement
	_, err := logapi.RegisterLoggerFactory(new(loggerzap.ZapFactoryRegister), multiOpts...)

	if err != nil {
		log.Println(err)
		return
	}

	// --- create logger factory manager

	logger := logapi.GetLogger("benshmark.test")

	//logger.Debug("debug message for  example", nil)
	//b.StartTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("info message for  example", nil)
		//logger.Warn("warn message for  example", nil)
	}
	//logger.Warn("warn message for  example", nil)
	//logger.Error("error  message for  example", nil)

}

//  Test_BasicCase1_Debug define bug info
func Benchmark_Zap_Factory_case1_advanced_structbean_example(b *testing.B) {
	//b.StopTimer()
	var multiOpts = make([]logapi.Option, 0)
	// --- construct layout ---
	var jsonLayout = zaplayout.NewJsonLayout()
	// --- set appender
	var consoleAppender = appender.NewConsoleAppender(jsonLayout)

	var loggerOpt1 = loggerzap.NewLoggerOption()
	loggerOpt1.SetLevel("info")
	loggerOpt1.AddAppender(consoleAppender)
	multiOpts = append(multiOpts, loggerOpt1)

	// use new or struct binding
	// create instance from implement
	_, err := logapi.RegisterLoggerFactory(new(loggerzap.ZapFactoryRegister), multiOpts...)

	if err != nil {
		log.Println(err)
		return
	}

	// --- create logger factory manager

	logger := logapi.GetLogger("benshmark.test")

	//logger.Debug("debug message for  example", nil)
	//b.StartTimer()
	for i := 0; i < b.N; i++ {

		var loggerBean = logapi.NewStructBean()

		loggerBean.LogString("testStringfield2", "logstring filed")
		loggerBean.LogBool("testBoolField2", false)
		loggerBean.LogFloat32("testfloat32Field2", 32.32)

		logger.Info("info message for  example", loggerBean)
		//logger.Warn("warn message for  example", nil)
	}
	//logger.Warn("warn message for  example", nil)
	//logger.Error("error  message for  example", nil)

}
