package log

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
	"os"
)

var logger *logrus.Logger


func init(){
	logger = &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Hooks: make(logrus.LevelHooks),
		Formatter: &prefixed.TextFormatter{
			DisableColors: true,
			TimestampFormat : "2006-01-02 15:04:05",
			FullTimestamp:true,
			ForceFormatting: true,
		},
	}
}

func SetLogToFile(fileName string){
	logger.Hooks.Add(NewFileHook(fileName))
}

func Info(ctx context.Context,args...interface{}){
	logger.Info(args)
}

func Infof(ctx context.Context,format string,args...interface{}){
	logger.Infof(format,args)
}

func Error(ctx context.Context,args...interface{}){
	logger.Error(args)
}

func Errorf(ctx context.Context,format string,args...interface{}){
	logger.Error(format,args)
}
