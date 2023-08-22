package logger

import (
	"github.com/sirupsen/logrus"
	"integrator_template/config"
	"os"
)

type fileHook struct {
	LevelsArr []logrus.Level
	Files     map[logrus.Level]*os.File
}

func (hook *fileHook) Fire(entry *logrus.Entry) error {
	for _, level := range hook.LevelsArr {
		if entry.Level <= level {
			entry.Logger.Out = hook.Files[level]
			break
		}
	}
	return nil
}

func (hook *fileHook) Levels() []logrus.Level {
	return hook.LevelsArr
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{e}
}

func InitLog(cfg *config.Config) {
	logger := logrus.New()
	logger.SetReportCaller(true)

	logger.SetFormatter(&logrus.JSONFormatter{})

	file, err := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatal(err)
	}
	logger.SetOutput(file)
	/*	logger.AddHook(&fileHook{
		LevelsArr: []logrus.Level{
			logrus.DebugLevel,
			logrus.InfoLevel,
			logrus.ErrorLevel,
		},
		Files: map[logrus.Level]*os.File{
			logrus.DebugLevel: file,
			logrus.InfoLevel:  file,
			logrus.ErrorLevel: file,
		},
	})*/

	e = logger.WithFields(logrus.Fields{
		"service": cfg.Server.Name,
	})
}
