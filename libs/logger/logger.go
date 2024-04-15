package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger = &logrus.Logger{
	Out: os.Stderr,
	Formatter: &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "Mon Jan 2 15:04:05 MST 2006",
	},
	Hooks: make(logrus.LevelHooks),
	Level: logrus.DebugLevel,
}

func Debug(args ...interface{}) {
	Logger.Debug(args...)
}
func DebugFn(fn logrus.LogFunction) {
	Logger.DebugFn(fn)
}
func Debugf(format string, args ...interface{}) {
	Logger.Debugf(format, args...)
}
func Debugln(args ...interface{}) {
	Logger.Debugln(args...)
}
func Error(args ...interface{}) {
	Logger.Error(args...)
}
func ErrorFn(fn logrus.LogFunction) {
	Logger.ErrorFn(fn)
}
func Errorf(format string, args ...interface{}) {
	Logger.Errorf(format, args...)
}
func Errorln(args ...interface{}) {
	Logger.Errorln(args...)
}
func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}
func FatalFn(fn logrus.LogFunction) {
	Logger.FatalFn(fn)
}
func Fatalf(format string, args ...interface{}) {
	Logger.Fatalf(format, args...)
}
func Fatalln(args ...interface{}) {
	Logger.Fatalln(args...)
}
func Info(args ...interface{}) {
	Logger.Info(args...) 
}
func Infofn(fn logrus.LogFunction) {
	Logger.InfoFn(fn)
}
func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args...)
}
func Infoln(args ...interface{}) {
	Logger.Infoln(args...)
}
func Log(level logrus.Level, args ...interface{}) {
	Logger.Log(level, args...)
}
func LogFn(level logrus.Level, fn logrus.LogFunction) {
	Logger.LogFn(level, fn)
}
func Logf(level logrus.Level, format string, args ...interface{}) {
	Logger.Logf(level, format, args...)
}
func Logln(level logrus.Level, args ...interface{}) {
	Logger.Logln(level, args...)
}
func Panic(args ...interface{}) {
	Logger.Panic(args...)
}
func PanicFn(fn logrus.LogFunction) {
	Logger.PanicFn(fn)
}
func Panicf(format string, args ...interface{}) {
	Logger.Panicf(format, args...)
}
func Panicln(args ...interface{}) {
	Logger.Panicln(args...)
}
func Print(args ...interface{}) {
	Logger.Print(args...)
}
func PrintFn(fn logrus.LogFunction) {
	Logger.PrintFn(fn)
}
func Printf(format string, args ...interface{}) {
	Logger.Printf(format, args...)
}
func Println(args ...interface{}) {
	Logger.Println(args...)
}
func Warn(args ...interface{}) {
	Logger.Warn(args...)
}
func WarnFn(fn logrus.LogFunction) {
	Logger.WarnFn(fn)
}
func Warnf(format string, args ...interface{}) {
	Logger.Warnf(format, args...)
}
func Warning(args ...interface{}) {
	Logger.Warning(args...)
}
func WarningFn(fn logrus.LogFunction) {
	Logger.WarningFn(fn)
}
func Warningf(format string, args ...interface{}) {
	Logger.Warningf(format, args...)
}
func Warningln(args ...interface{}) {
	Logger.Warningln(args...)
}
func Warnln(args ...interface{}) {
	Logger.Warnln(args...)
}