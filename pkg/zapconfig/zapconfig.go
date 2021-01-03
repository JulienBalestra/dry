package zapconfig

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZapConfig() *zap.Config {
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.MessageKey = "m"
	zapConfig.EncoderConfig.LevelKey = "l"
	zapConfig.EncoderConfig.TimeKey = "ts"
	zapConfig.EncoderConfig.EncodeTime = epochTimeEncoder
	zapConfig.EncoderConfig.CallerKey = "c"
	return &zapConfig
}

func epochTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	secs := t.Second()
	enc.AppendInt(secs)
}
