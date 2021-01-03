package zapconfig

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapConfig() *zap.Config {
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig = zapcore.EncoderConfig{
		MessageKey: "m",
		LevelKey:   "l",
		TimeKey:    "ts",
		EncodeTime: epochTimeEncoder,
		CallerKey:  "c",
	}
	return &zapConfig
}

func epochTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	secs := t.Second()
	enc.AppendInt(secs)
}
