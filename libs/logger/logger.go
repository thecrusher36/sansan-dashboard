package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger = &logrus.Logger{
	Out: os.Stderr,
	Formatter: &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "Mon Jan 2 15:04:05 MST 2006",
	},
	Hooks: make(logrus.LevelHooks),
	Level: logrus.DebugLevel,
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}
func DebugFn(fn logrus.LogFunction) {
	logger.DebugFn(fn)
}
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}
func Debugln(args ...interface{}) {
	logger.Debugln(args...)
}
func Error(args ...interface{}) {
	logger.Error(args...)
}
func ErrorFn(fn logrus.LogFunction) {
	logger.ErrorFn(fn)
}
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}
func Errorln(args ...interface{}) {
	logger.Errorln(args...)
}
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}
func FatalFn(fn logrus.LogFunction) {
	logger.FatalFn(fn)
}
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}
func Fatalln(args ...interface{}) {
	logger.Fatalln(args...)
}
func Info(args ...interface{}) {
	logger.Info(args...) 
}
func Infofn(fn logrus.LogFunction) {
	logger.InfoFn(fn)
}
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}
func Infoln(args ...interface{}) {
	logger.Infoln(args...)
}
func Log(level logrus.Level, args ...interface{}) {
	logger.Log(level, args...)
}
func LogFn(level logrus.Level, fn logrus.LogFunction) {
	logger.LogFn(level, fn)
}
func Logf(level logrus.Level, format string, args ...interface{}) {
	logger.Logf(level, format, args...)
}
func Logln(level logrus.Level, args ...interface{}) {
	logger.Logln(level, args...)
}
func Panic(args ...interface{}) {
	logger.Panic(args...)
}
func PanicFn(fn logrus.LogFunction) {
	logger.PanicFn(fn)
}
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}
func Panicln(args ...interface{}) {
	logger.Panicln(args...)
}
func Print(args ...interface{}) {
	logger.Print(args...)
}
func PrintFn(fn logrus.LogFunction) {
	logger.PrintFn(fn)
}
func Printf(format string, args ...interface{}) {
	logger.Printf(format, args...)
}
func Println(args ...interface{}) {
	logger.Println(args...)
}
func Warn(args ...interface{}) {
	logger.Warn(args...)
}
func WarnFn(fn logrus.LogFunction) {
	logger.WarnFn(fn)
}
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}
func Warning(args ...interface{}) {
	logger.Warning(args...)
}
func WarningFn(fn logrus.LogFunction) {
	logger.WarningFn(fn)
}
func Warningf(format string, args ...interface{}) {
	logger.Warningf(format, args...)
}
func Warningln(args ...interface{}) {
	logger.Warningln(args...)
}
func Warnln(args ...interface{}) {
	logger.Warnln(args...)
}