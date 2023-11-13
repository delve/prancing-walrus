package log

import (
	"go.uber.org/zap"
)

var (
	Logger     *zap.SugaredLogger
	FastLogger *zap.Logger
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

func Fatalw(msg string, keysAndValues ...interface{}) {
	Logger.Fatalw(msg, keysAndValues...)
}

func getLogger() (err error) {
	err = nil
	// TODO: NewProduction is a canned set of production-ready configs for the logger.
	//       expand to customizable configs and enable changing log level. from env var? updating via chat cmds?
	FastLogger, err = zap.NewProduction()
	if err != nil {
		return
	}
	Logger = FastLogger.Sugar()
	return
}
