package log

import (
	"go.uber.org/zap"
)

var (
	Logger     *zap.SugaredLogger
	FastLogger *zap.Logger
	cfg        zap.Config
)

func init() {
	err := getLogger()
	if err != nil {
		panic(err)
	}
}

func Infow(msg string, keysAndValues ...interface{}) {
	Logger.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	Logger.Warnw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	Logger.Fatalw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	Logger.Errorw(msg, keysAndValues...)
}

func getLogger() (err error) {
	err = nil
	cfg = zap.NewProductionConfig()

	FastLogger = zap.Must(cfg.Build())

	Logger = FastLogger.Sugar()
	return
}

func SetLevelDebug() {
	cfg.Level.SetLevel(zap.DebugLevel)
	Infow("updated log level", "level", "debug")
}
