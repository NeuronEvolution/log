package log

import (
	"go.uber.org/zap"
	"log"
	"os"
	"reflect"
)

type SystemLog struct {
	logger *zap.Logger
}

func (l *SystemLog) Write(p []byte) (n int, err error) {
	l.logger.Info(string(p))
	return len(p), nil
}

func Init() {
	var l *zap.Logger
	var err error

	if os.Getenv("DEBUG") != "" {
		l, err = zap.NewDevelopment(zap.AddCaller())
	} else {
		l, err = zap.NewProduction(zap.AddCaller())
	}

	if err != nil {
		log.Fatal(err)
		return
	}

	zap.ReplaceGlobals(l)

	log.SetOutput(&SystemLog{logger: zap.L().WithOptions(zap.AddCallerSkip(3)).Named("golog")})
}

func TypedLogger(i interface{}) *zap.Logger {
	return zap.L().Named(reflect.TypeOf(i).String())
}
