package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

type LogrusFileHook struct {
	file      *os.File
	formatter *logrus.JSONFormatter
}

func NewFileHook(fileName string) *LogrusFileHook {
	jsonFormat := &logrus.JSONFormatter{}
	file, err := os.Create(fileName)
	if err != nil {
		panic("Cannot open file for logging")
	}

	return &LogrusFileHook{
		file:      file,
		formatter: jsonFormat,
	}
}

// Fire event
func (hook *LogrusFileHook) Fire(entry *logrus.Entry) error {

	plainformat, err := hook.formatter.Format(entry)
	line := string(plainformat)
	_, err = hook.file.WriteString(line)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to write file on filehook(entry.String)%v", err)
		return err
	}

	return nil
}

func (hook *LogrusFileHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}
