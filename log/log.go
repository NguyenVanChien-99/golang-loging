package log

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
	"os"
)

var logger *logrus.Logger

const MIDDLEWARE_REQUEST_UID = "X-Request-ID"

func init() {
	logger = &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Hooks: make(logrus.LevelHooks),
		Formatter: &prefixed.TextFormatter{
			DisableColors:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}
}

func SetLogToFile(fileName string) {
	logger.Hooks.Add(NewFileHook(fileName))
}

func Info(ctx context.Context, args ...interface{}) {
	getRequestUid(ctx).Info(args)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	getRequestUid(ctx).Infof(format, args)
}

func Error(ctx context.Context, args ...interface{}) {
	getRequestUid(ctx).Error(args)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	getRequestUid(ctx).Errorf(format, args)
}

func getRequestUid(ctx context.Context) *logrus.Entry {
	requestUid := ctx.Value(MIDDLEWARE_REQUEST_UID)
	if requestUid == nil || fmt.Sprintf("%s", requestUid) == "" {
		return logger.WithContext(ctx)
	}
	return logger.WithField("req_uid", requestUid)
}
