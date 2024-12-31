package ezgo

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Scope struct {
	parent *Scope
	// properties of this scope that is immutable
	logger *zap.Logger
}

func NewScope(
	logger *zap.Logger,
) *Scope {
	return &Scope{
		logger: logger,
	}
}

func NewScopeWithDefaultLogger() *Scope {
	return NewScope(createDefaultLogger())
}

func createDefaultLogger() *zap.Logger {
	zap.NewProduction()
	location, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}

	// Custom Time Encoder for PST
	pstTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.In(location).Format("2006-01-02 15:04:05"))
	}

	// Custom encoder config
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:          "ts",
		LevelKey:         "level",
		MessageKey:       "msg",
		NameKey:          "logger",
		StacktraceKey:    "stacktrace",
		EncodeTime:       pstTimeEncoder, // Use the custom PST time encoder
		EncodeLevel:      zapcore.CapitalLevelEncoder,
		EncodeDuration:   zapcore.StringDurationEncoder,
		ConsoleSeparator: " | ",
	}

	loc, _ := time.LoadLocation("America/Los_Angeles")
	logFilePath := fmt.Sprintf("logs/%s.txt", time.Now().In(loc).Format("2006-01-02_15-04"))
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		//		Encoding:         "console",
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout", logFilePath},
		ErrorOutputPaths: []string{"stderr", logFilePath},
	}

	return Must(config.Build())
}

func (s *Scope) GetLogger() *zap.Logger {
	return findFirstNonNilPropoerty(s, "logger", func(s *Scope) *zap.Logger {
		return s.logger
	})
}

func (s *Scope) Close() {
	for ; s != nil; s = s.parent {
		s.logger.Sync()
	}
}

func (s *Scope) WithLogger(logger *zap.Logger) *Scope {
	return &Scope{
		parent: s,
		logger: logger,
	}
}

func findFirstNonNilPropoerty[T any](s *Scope, propertyName string, getProperty func(*Scope) *T) *T {
	for ; s != nil; s = s.parent {
		if p := getProperty(s); p != nil {
			return p
		}
	}
	AssertAlwaysf("Nil property %s", propertyName)
	return nil
}
