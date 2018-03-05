package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"reflect"
	"time"
)

type SystemLog struct {
	logger *zap.Logger
}

func (l *SystemLog) Write(p []byte) (n int, err error) {
	l.logger.Info(string(p))
	return len(p), nil
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02T15:04:05.000"))
}

func Init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = timeEncoder
	l, err := config.Build(zap.AddCaller())
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
