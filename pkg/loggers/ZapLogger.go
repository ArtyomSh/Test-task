package loggers

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
)

func New(verbosity bool) *zap.Logger {

	var level string
	if verbosity {
		level = "debug"
	} else {
		level = "warn"
	}

	cfgJSON := []byte(fmt.Sprintf(`{
   "level": "%s",
   "encoding": "json",
   "outputPaths": ["stdout"],
   "errorOutputPaths": ["stderr"],
   "encoderConfig": {
     "messageKey": "message",
     "levelKey": "level",
     "levelEncoder": "lowercase"
   }
 }`, level))

	var cfg zap.Config
	if err := json.Unmarshal(cfgJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	return logger
}
