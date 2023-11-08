package logs

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	config := zap.NewDevelopmentConfig()

	// Custom time encoder to format the time as dd/mm/yy HH:mm
	config.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("02/01/06 15:04")) // dd/mm/yy HH:mm format
	}

	// Custom level encoder to add color
	config.EncoderConfig.EncodeLevel = func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		var color string
		switch level {
		case zapcore.DebugLevel:
			color = "\033[36m" // Cyan for debug
		case zapcore.InfoLevel:
			color = "\033[32m" // Green for info
		case zapcore.WarnLevel:
			color = "\033[33m" // Yellow for warn
		case zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
			color = "\033[31m" // Red for error and above
		default:
			color = "\033[0m" // Default
		}
		enc.AppendString(color + level.CapitalString() + "\033[0m")
	}

	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message interface{}, fields ...zap.Field) {
	var errMessage string
	switch v := message.(type) {
	case error:
		errMessage = v.Error()
	case string:
		errMessage = v
	default:
		errMessage = "Unknown error type"
		fields = append(fields, zap.Any("error", message))
	}
	log.Error(errMessage, fields...)
}
